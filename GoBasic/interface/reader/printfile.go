package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printFile(filename string)  {
	file, err:=os.Open(filename)
	if err!=nil{
		panic(err)
	}

	printFileContents(file) //File 类型实现了 Reader 接口
}

func printFileContents(reader io.Reader)  {
	scanner :=bufio.NewScanner(reader)
	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
}


func main() {
	printFile("abc.txt")
	fmt.Println("=============")
	s:=`abc"d"
	efg`
	printFileContents(strings.NewReader(s))
}
