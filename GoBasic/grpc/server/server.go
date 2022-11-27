package main

import (
	proto "grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	// 监听8000 端口, 返回一个 listener 和 error
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}

	s:= proto.Server{}

	grpcServer := grpc.NewServer()

	//注册一个server
	proto.RegisterChatServiceServer(grpcServer,&s)

	// server 开始监听
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fail to serve: %v", err)
	}


}
