package main

import "fmt"

func getSum() func()   {
	sum:=0
	return func()   {
		sum+=1
		fmt.Println(sum)
	}
}

func main() {
	f:=getSum()
	f()
	f()
	f()
	f()
}

