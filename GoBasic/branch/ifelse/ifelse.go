package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const  filename string="abc.txt"
	contents, err:=ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s \n", contents)
	}
}
