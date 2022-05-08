package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生成一个 chan int 并 return 出去, 这个 chan int 被不停地修改
func generator() chan int  {
	out:= make(chan int)
	go func() {
		i:=0
		for{
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			out <- i //往channel 里面传消息时, main routine 也会被阻塞
			i++
		}
	}()
	return out
}

//一个读取 channel int 变量的函数
func worker(id int, c chan  int)  {
	go func() {
		for{
			time.Sleep(time.Second)//故意休眠长时间,让worker接收另一个channel的时间推后
			n,ok:= <-c
			if !ok{
				break
			}
			fmt.Printf("Worker %d received %d \n", id, n)
		}
	}()
}

// 一个创建 channel int 并不断读取 channel int 的函数
func createWorker(id int ) chan<- int {
	c := make(chan int)
	worker(id,c)
	return  c
}

func main() {

	var c1, c2= generator(),generator()
	var worker =createWorker(0) // worker 是一个 chan int
	var values[] int


	tm:=time.After(10*time.Second)
	tick:=time.Tick(time.Second)//每隔一秒
	for  {
		var activeWorker chan<- int // 为 nil
		var activeValue int
		if len(values)>0 {
			activeWorker=worker
			activeValue=values[0]
		}
		// select 相当于针对 channel 的 switch, 每一个 case 执行完之前都会等待

		select {//n 去接收c1, c2, 谁来的快收谁
		case n:= <- c1:
			values=append(values,n)
		case n:=<- c2:
			values=append(values,n)
		case <-tick:
			fmt.Println("queue len: ", len(values))
		//如果两次chan int 的生成时间超过了800mS
		case <-time.After(800*time.Millisecond):
			fmt.Println("timeout")

		case  activeWorker<-activeValue: //若 activateWorker 为 nil 则不满足这个 case, 进入下一个循环,在前面两个case循环
			values=values[1:]
		case <-tm:
			fmt.Println("bye bye")
			return

		//这里不能加default, 如果加了, generator 里面的routine往channel里面传递的时候虽然会阻塞,
		// 但是main routine 这边还是极大概率会运行到default 而 generator routine里面大多数时间都在sleep,导致channel为nil
		//default:
		//	fmt.Println("No value received from c1,c2")
		}
	}

}
