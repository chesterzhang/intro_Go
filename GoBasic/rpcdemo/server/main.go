package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpcdemo"
)


func main() {
	rpc.Register(rpcdemo.DemoService{}) //注册一个 service

	listener,err:=net.Listen("tcp",":1234")// 监听这个 端口


	if err!=nil {
		panic(err)
	}

	for  {
		conn,err:= listener.Accept() // 接收传入的 连接
		if err!=nil{
			log.Printf("Accept erro :%v",err)
			continue
		}
		go jsonrpc.ServeConn(conn) // 处理任务
	}


}
