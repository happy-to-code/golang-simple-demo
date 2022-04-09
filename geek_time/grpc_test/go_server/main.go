// main.go
package main

import (
	"GoProjectDemo/geek_time/grpc_test/go_server/controller/hello_controller"
	"GoProjectDemo/geek_time/grpc_test/go_server/proto/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

const Address = "0.0.0.0:9090"

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 服务注册
	hello.RegisterHelloServer(s, &hello_controller.HelloController{})
	log.Println("Listen on " + Address)
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
