package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock sync.Mutex //锁
}

func (a *atomicInt) increment(){
	fmt.Println("safe increment")
	//如果想要对一个代码块进行保护, 则用一个匿名函数
	func(){
		a.lock.Lock()
		defer a.lock.Unlock() //在函数执行完之前释放锁
		a.value++
	}()

}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer  a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
