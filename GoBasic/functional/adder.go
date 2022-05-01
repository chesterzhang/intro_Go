package main

import "fmt"

// 函数式编程, 函数里面有一个局部变量 sum
//adder 函数, 没有参数, 返回值是一个函数
func adder() func(int) int {
	sum:=0
	return func(v int) int {
		sum+=v
		return  sum
	}
}

func main() {
	a:=adder()
	for i:=0;i<10 ;	i++  {
		fmt.Printf("0+...+%d = %d \n",i,a(i))
	}
}
