package main

import (
	"fmt"
)

func tryRecover()  {
	defer func() {
		r:= recover() // 返回一个interface{}类型, 也就是 任何类型
		// recover 作用
		//1. 仅在 defer 中调用
		//2. 获取panic 值
		err,ok:=r.(error)//如果确实是 error 类型
		if ok{
			fmt.Println("Error occurred:", err)
		}else{
			panic(r)
		}
	}()
	//panic(errors.New("This is an error"))
	// panic 作用
	//1. 停止当前函数执行
	//2. 一直向上返回, 执行每一层的 defer
	//如果没有遇见 recover, 程序退出

	b:=0
	a:=5/b
	fmt.Println(a)
}

func main() {
	tryRecover()
}
