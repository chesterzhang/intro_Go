package main

import (
	"fmt"
	"runtime"
	"time"
)

// goroutine 可能的切换点
// I/O, select, channerl, 等待锁, 函数调用, runtime.Gosched()

func main() {
	var a [10]int

	//协程是一种轻量级线程, 与线程的主要区别在于协程是非抢占式的, 需要协程自己主动交出控制权
	//开10个协程, 每一个协程都去做 a[i]++(每一个协程的i不同)
	for i:=0;i<10 ; i++  {
		 go func(i int){
		 	for{
		 		a[i]++
		 		runtime.Gosched()//让出对 a[]的 控制权, 因为协程是不能被抢占的, 只能主动让出
			}
		 }(i)//若不把i 传进去则 内部的匿名函数中的i是外面for 循环的i, 可以取到10,越界
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
