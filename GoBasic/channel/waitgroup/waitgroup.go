package main

import (
	"fmt"
	"sync"
)

// 一个创建 worker 并不断读取 worker.in , 并且返回 worker 的函数
func createWorker(id int, wg *sync.WaitGroup ) worker {
	w:=worker{
		in:  make(chan  int),
		wg: wg,
	}

	doWorker(id,w.in,wg)
	return  w
}

//一个读取 channel int 变量的函数, 读取完以后将 channel bool 设为true
func doWorker(id int, c chan  int, wg *sync.WaitGroup )  {
	go func() {//goroutine 去读 chan int, 并设置  channel bool
		for{
			n,ok:= <-c
			if !ok{
				break
			}
			fmt.Printf("Worker %d received %c \n", id, n)
			wg.Done()
		}
	}()
}


type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func chanDemo()  {
	var workers [10]worker // 10个worker

	//wait group
	var wg sync.WaitGroup
	wg.Add(20)//等待20个任务

	for i:=0;i<10 ;i++  {
		workers[i]= createWorker(i,&wg)//将10个 channel int 分发给10个worker

	}


	for i,worker:=range workers  {
		worker.in <- 'a'+i //对channel int 进行改变, 同时 createWorker 里面在读
	}

	for i,worker:=range workers  {
		worker.in <- 'A'+i //对channel int 进行改变, 同时 createWorker 里面在读
	}

	wg.Wait()

}




func main() {
	chanDemo()

}

