package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer()  {
	defer fmt.Println(1) //在函数return之前打印, 如果有多个 defer, 则defer 语句进入一个栈, 先进后出
	defer fmt.Println(2)
	fmt.Println(3)
	return

}

func writeFile(filename string){
	//file,err:=os.Create(filename)
	file,err:=os.OpenFile(filename, os.O_EXCL | os.O_CREATE, 0666)//如果出错, 返回一个 pathError类型

	//err=errors.New("This is a custom error")//自定义 error
	if err!=nil{
		pathError, ok:=err.(*os.PathError)
		if !ok{//如果不是pathError
			panic(err)
		}else {//如果是pathError
			fmt.Println(pathError.Op, pathError.Path,pathError.Err)
		}
		return
	}

	defer  file.Close()

	writer:=bufio.NewWriter(file)
	defer  writer.Flush()

	for i := 0; i<11; i++  {
		fmt.Fprintln(writer,i)
	}

}

func main() {
	tryDefer()
	writeFile("zeroten.txt")
}
