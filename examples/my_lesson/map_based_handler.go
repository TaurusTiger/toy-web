package main

import (
	"net/http"
)

type HandlerBasedOnMap struct {
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	key := h.Key(method, pattern)
	h.handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.Key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not found"))
	}
}

func (h *HandlerBasedOnMap) Key(method string, pattern string) string {
	return method + "#" + pattern
}

type Routable interface {
	Route(method string, pattern string, handleFunc func(ctx *Context))
}

type HandlerInterface interface {
	http.Handler
	// Route(method string, pattern string, handleFunc func(ctx *Context))
	Routable
}

// 一种常用的Go设计模式
// 用于确保HandlerBasedOnMap肯定实现了这个接口
// 起到保护作用
var _ HandlerInterface = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() HandlerInterface {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context)),
	}
}

//type teset struct{}
//
//func (t *teset) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (t *teset) Route(method string, pattern string, handleFunc func(ctx *Context)) {
//	//TODO implement me
//	panic("implement me")
//}


