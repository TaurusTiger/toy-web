package main

import (
	"fmt"
	"net/http"
)

func home2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

//func main() {
//	server := NewHttpServer("test-server")
//
//	server.Route("/", home)
//	server.Route("/user/signup", SignUp)
//
//	log.Fatal(server.Start(":8800"))
//}
