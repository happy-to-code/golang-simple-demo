package main

import (
	"GoProjectDemo/geek_time/grpc_test/go_client/proto/hello"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
)

const Address = "0.0.0.0:9090"

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	// 一定执行
	defer conn.Close()
	// 初始化客户端
	c := hello.NewHelloClient(conn)
	// 调用SayHello方法
	res, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "--Hello World==="})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("------------>", res.Message)
	// 调用LotsOfReplies方法
	stream, err := c.LotsOfReplies(context.Background(), &hello.HelloRequest{Name: "Hello World"})
	if err != nil {
		log.Fatalln(err)
	}
	for true {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream.Recv: %v", err)
		}
		log.Printf("%s", res.Message)
	}
}
