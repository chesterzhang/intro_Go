package main

import (
	"context"
	"fmt"
	"time"
)

func watch(ctx context.Context, name string)  {
	//这里必须要 go func, 否则一直在这个 for 循环中, main 函数执行不到 cancel 调用
	go func() {
		for  {
			select {
			case <-ctx.Done():
				fmt.Println(name+" stop")
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}()
}

func main() {
	rootCtx:=context.Background()
	cancelCtx, cancel:=context.WithCancel(rootCtx)
	watch(cancelCtx,"cancelCtx")
	time.Sleep(2*time.Second)
	cancel()
	time.Sleep(2*time.Second)

}
