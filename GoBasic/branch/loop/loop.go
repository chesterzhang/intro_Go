package main

import (
	"bufio"
	"fmt"
	"os"
)

func printFile(filename string)  {
	file, err:=os.Open(filename)
	if err!=nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for ; scanner.Scan();  {
		fmt.Println(scanner.Text())
	}

}

func add()  {
	var i,sum int=1,0
	for ;i<5 ; i++  {
		sum+=i
	}
	fmt.Println(sum)
}
func main() {
	add()
	printFile("abc.txt")
}
