package main

import "fmt"

// 数组是值类型, 不是引用类型
func testArray(arr [5]int)  {
	arr[0]=100
}

func main() {
	var arr1 [5]int
	var arr2 =[5]int{1,3,5,7,9}
	var arr3 [2][2]int
	fmt.Println(arr1,arr2,arr3)

	//遍历数组
	for  i:= 0; i<len(arr2) ;i++  {
		fmt.Println(arr2[i])
	}

	for i:=range arr2{
		fmt.Println(arr2[i])
	}

	for i,v :=range arr2{
		fmt.Println(i,v)
	}

	testArray(arr1)
	fmt.Println(arr1)
}
