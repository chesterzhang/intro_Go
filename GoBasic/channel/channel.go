package main

import (
	"fmt"
	"time"
)

// 一个创建 channel int 并不断读取 channel int 的函数
func createWorker(id int ) chan<- int {
	c := make(chan int)
	//go func() {
	//	for{
	//		fmt.Printf("Worker %d received %c \n", id, <-c)
	//	}
	//}()
	worker(id,c)
	return  c
}

func chanDemo()  {
	var channels [10]chan<- int // chan<- int 表示 这个 chan int 变量只能被写入, 不能被读取

	for i:=0;i<10 ;i++  {
		channels[i]= createWorker(i)//将10个 channel int 分发给10个worker
	}

	for i:=0;i<10 ;i++   {
		channels[i] <- 'a'+i //对channel int 进行改变, 同时 createWorker 里面在读
		//n:= channels[i] //报错, 因为 chan<- int 只能被写, 不能被读
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel()  {
	c:=make(chan  int,3) //可以往 chan int 发3个数据,不读也不会报错
	go worker(0,c) //l 个 用来读 channel int 的函数
	c <-'A'
	c <-'B'
	c <-'C'
	c <-'D'
	time.Sleep(time.Millisecond)
}

//一个读取 channel int 变量的函数
func worker(id int, c chan  int)  {
	go func() {
		for{
			n,ok:= <-c
			if !ok{
				break
			}
			fmt.Printf("Worker %d received %c \n", id, n)
		}
	}()
}

func channelClose()  {
	c:=make(chan  int, 3)
	go worker(0,c)
	c <-'A'
	c <-'B'
	c <-'C'
	c <-'D'
	close(c)//关闭这个channel, 读取 channel int 的地方还会一直读 1 mS
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	//bufferedChannel()
	//channelClose()
}
