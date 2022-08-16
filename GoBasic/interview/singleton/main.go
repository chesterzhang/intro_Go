package main

import (
	"fmt"
	"time"
)

type singleton struct {
	t time.Time
}



var instance *singleton

func  getInstance() *singleton  {
	if instance==nil {
		instance=&singleton{t:time.Now()}
		return instance
	}else {
		return instance
	}
}

func main() {
	a:=getInstance()
	fmt.Println(a.t)
	time.Sleep(time.Second)
	b:=a
	fmt.Println(b.t)

}
