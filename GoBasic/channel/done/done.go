package main

import (
	"fmt"
)

// 一个创建 worker 并不断读取 worker.in , 并且返回 worker 的函数
func createWorker(id int ) worker {
	w:=worker{
		in:  make(chan  int),
		done: make(chan  bool),
	}

	doWorker(id,w.in,w.done)
	return  w
}

//一个读取 channel int 变量的函数, 读取完以后将 channel bool 设为true
func doWorker(id int, c chan  int, done chan bool)  {
	go func() {//goroutine 去读 chan int, 并设置  channel bool
		for{
			n,ok:= <-c
			if !ok{
				break
			}
			fmt.Printf("Worker %d received %c \n", id, n)

			done<- true //读取channel int完成, 将 done设为 true, 相当于用一个channel去通知外面已经完成任务了

		}
	}()
}


type worker struct {
	in chan int
	done chan bool
}

func chanDemo()  {
	var workers [10]worker // 10个worker

	for i:=0;i<10 ;i++  {
		workers[i]= createWorker(i)//将10个 channel int 分发给10个worker

	}

	for i,worker:=range workers  {
		worker.in <- 'a'+i //对channel int 进行改变, 同时 createWorker 里面在读
	}

	for _,worker := range workers{
		<-worker.done//只有当 wokers[i].done 里面有值的时候,这里才能读到,才能结束, 所以不再需要等待1mS
	}

	for i,worker:=range workers  {
		worker.in <- 'A'+i //对channel int 进行改变, 同时 createWorker 里面在读
	}
	for _,worker := range workers{
		<-worker.done//只有当 wokers[i].done 里面有值的时候,这里才能读到,才能结束, 所以不再需要等待1mS
	}

}




func main() {
	chanDemo()

}
