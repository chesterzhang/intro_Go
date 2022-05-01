package  main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a,b int, op string) (int,error) {
	switch op {
	
	case "+":
		return  a+b, nil
	case "-":
		return a-b, nil
	case "*":
		return a*b, nil
	case "/":
		q,_ :=div(a,b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsorported operation : %s", op)
	}
}

//整除取余
func div(a,b int) (int,int){
	return a/b, a%b
}

func applyFun(op func(int,int) int,a ,b int) int  {
	p:= reflect.ValueOf(op).Pointer()// 拿到需要调用函数的指针
	opName:=runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args : %d %d \n", opName,a,b)
	return op(a,b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(eval(6,3,"*"))
	fmt.Println(div(7,3))
	fmt.Println(eval(7,3,"/"))
	fmt.Println(eval(7,3,"&"))
	fmt.Println(applyFun(pow,2,4))
}
