package main

import (
	"fmt"
	"sync"
)

func main() {

	c:=make(chan  int)
	wg:=sync.WaitGroup{}
	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c<-i
		}(i)
		fmt.Println(<-c)
	}

	wg.Wait()
}
