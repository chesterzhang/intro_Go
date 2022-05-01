package main

import (
	"fmt"
)

func grade(score int) string  {
	var g string
	switch  {
	default:
		panic(fmt.Sprintf("Wrong score : %d \n",score)) //中断报错
	case score<60:
		g="F"
	case score<80:
		g="C"
	case score<90:
		g="B"
	case score<=100:
		g="A"
	}
	return g
}

func main() {
	fmt.Println(grade(10),grade(60),grade(80),grade(90),grade(100),grade(200),grade(-1))

}

