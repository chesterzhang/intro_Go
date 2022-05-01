package main

import (
	"fmt"
	"queue"
)

func main() {
	var q =queue.Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q)
	q.Pop()
	q.Pop()
	fmt.Println(q.IsEmpty())
}
