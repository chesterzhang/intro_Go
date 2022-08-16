package main

import (
	"context"
	"fmt"
)

func A(ctx context.Context)  {
	fmt.Println("This is function A")
	ctxB:=context.WithValue(ctx,"k1","v1")
	B(ctxB)
}

func B(ctx context.Context)  {
	fmt.Println("This is function B")
	ctxC:=context.WithValue(ctx,"k2","v2")
	C(ctxC)
}

func C(ctx context.Context)  {
	fmt.Println("This is function C")
	fmt.Println(ctx.Value("k1"))
	fmt.Println(ctx.Value("k2"))
}

func main() {
	rootCtx:=context.Background()

	A(rootCtx)
}
