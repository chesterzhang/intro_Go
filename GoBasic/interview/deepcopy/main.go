package main

import "fmt"

func main() {

	a:=[]int{1,2,3}
	b:=make([]int,3)
	copy(b,a)
	b[0]=-1
	fmt.Println(a)
	fmt.Println(b)

	c:=a
	fmt.Println(a)
	fmt.Println(c)


}
