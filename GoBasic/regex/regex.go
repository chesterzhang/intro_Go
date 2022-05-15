package main

import (
	"fmt"
	"regexp"
)

const text  =` My email is tom@gmail.com@@@@  
email2 is john@qq.com 
email3 is Lily@xxu.edu.hk`

func main() {

	re,err:=regexp.Compile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)//.+表示一个或多个字母, \. 必须匹配 . ,()提取
	if err!=nil {
		panic(err)
	}
	//match:=re.FindString(text)
	//match:=re.FindAllString(text,-1)
	match:=re.FindAllStringSubmatch(text,-1)

	for _,m:=range match{
		fmt.Println(m)
	}
	fmt.Println(match)
}
