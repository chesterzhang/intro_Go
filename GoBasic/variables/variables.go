package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// go 内建变量类型
// bool
// string
// (u)int (u)int8 (u)int16 (u)int32 (u)int64 uintptr
// flaot32 float64
// byte rune 字符型
// complex64 complex128

var (
	aa int=10
	ss string="ss"
	sss=string("sss")//string 的另一种初始化
	bb bool=false

)


func variableZeroValue()  {
	var a int
	var s string
	fmt.Printf("%d %s\n",a,s)
	fmt.Printf("%d %q\n",a,s)
}

func variableInit()  {
	var a, b  int=1,3
	var c  = bool(true)
	var s string="abc"
	fmt.Println(a,b,c,s)
}

func variableTypeDeduction()  {
	var a, b,c  =1,3,true
	var s string="abc"
	fmt.Println(a,b,c,s)
}

func complexNum()  {
	var c complex128=3+4i
	fmt.Println(cmplx.Abs(c))
}

func euler()  {
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi)) //欧拉公式
	fmt.Println(cmplx.Exp(1i*math.Pi)) //欧拉公式
}

func triangle(){
	var a,b int=3,4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}

func consts()  {
	const filename string = "abc.txt"
	const a,b  = 3,4
	fmt.Println(filename,a,b)
}


func enums(){
	const(
		cpp= iota
		_
		java
		python
		golang
	)

	const(
		b= 1<< (10*iota)
		kb
		mb
		gb
		tb
	)
	fmt.Println(cpp,java,python,golang)
	fmt.Println(b,kb,mb,gb,tb)
}

func main() {
	fmt.Println(sss)
	variableZeroValue()
	variableInit()
	variableTypeDeduction()
	fmt.Println(aa,ss,bb)
	complexNum()
	euler()
	triangle()
	consts()
	enums()
}
