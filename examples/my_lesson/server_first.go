package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
}

// Route2 这里就比较抽象处理
func (s *sdkHttpServer) Route2(pattern string, handleFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		// 下面这段直接重新封装到了NewContext里了
		//ctx := &Context{
		//	R: request,
		//	W: writer,
		//}
		ctx := NewContext(writer, request)
		handleFunc(ctx)
	})
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{R: request, W: writer}
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

type Hander map[string][]string

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	} // 这里的返回必须是指针，是因为implement接口的时候，接收器是指针
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Fprintf(w, "read body failed: %v", err)
	//	// 要返回掉，不然就会继续执行后面的代码
	//	return
	//}
	//err = json.Unmarshal(body, req)
	////if err != nil {
	////	fmt.Fprintf(w, "deserialized failed: %v", err)
	////	// 要返回掉，不然就会继续执行后面的代码
	////	return
	////}

	ctx := &Context{
		W: w,
		R: r,
	}

	err := ctx.ReadJson(req)

	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
	}

	//
	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败：%v", err)
	}

	//respJson, errr := json.Marshal(resp)
	// 返回一个虚拟的 user id 表示注册成功了
	//fmt.Fprintf(w, "%d", err)
	//fmt.Fprintf(w, string(respJson))
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func main() {
	server := NewHttpServer("test-server")

	server.Route("/", home)
	server.Route("/user/signup", SignUp)

	log.Fatal(server.Start(":8800"))
}
