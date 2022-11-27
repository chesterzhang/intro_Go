package main

import (
	proto "grpc/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {

	//获得一个 Client 连接
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// 获得一个 client
	c := proto.NewChatServiceClient(conn)

	// grpc 调用远程的 SayHello
	response, err := c.SayHello(context.Background(), &proto.ChatMessage{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}
