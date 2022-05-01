package main

import "fmt"

func swap1(a, b int)  {
	b,a =a,b
}

func swap2(a, b *int)  {
	*b,*a =*a,*b
}

func main() {
	var a,b int =3,4
	swap1(a,b)
	fmt.Println(a,b)

	swap2(&a,&b)
	fmt.Println(a,b)
}
