package main

import (
	"fmt"
	"sync"
)

func main() {

	c1:=make(chan int,1)
	c2:=make(chan int)

	wg:=sync.WaitGroup{}
	wg.Add(2)

	c1<-1
	go func() {
		defer  wg.Done()
		for i:=0; i<10;i++{
			<-c1
			fmt.Println("A")
			c2<-1
		}
	}()

	go func() {
		defer  wg.Done()
		for i:=0; i<10;i++{
			<-c2
			fmt.Println("B")
			c1<-1
		}
	}()
	wg.Wait()
}
