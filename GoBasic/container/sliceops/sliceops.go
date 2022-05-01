package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf(" %v,len=%d, cap=%d \n", s,len(s), cap(s))
}

func main() {
	var s []int
	for i := 0; i<10;i++  {
		printSlice(s)
		s=append(s,i) //每当空间不够的时候,就会扩充两倍空间
	}
	fmt.Println("s=",s)

	s1:=make([]int,10)
	s2:=make([]int,10,32)
	printSlice(s1)
	printSlice(s2)

	fmt.Println("Copying slice")
	copy(s2,s) //将s copy到s2中去
	printSlice(s2)

	fmt.Println("Deleting slice") //删掉 s2 中的 3
	s2=append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	s2 =s2[1:]
	printSlice(s2)

	fmt.Println("Popping from back")
	s2 =s2[:len(s2)-1]
	printSlice(s2)
}
