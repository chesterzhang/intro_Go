package main

import "fmt"

func updataSlice(s  []int)  {
	s[0]=100
}



func main() {
	arr:=[...]int{0,1,2,3,4,5,6,7} // [...] 表示不需要指定长度, 而是根据初始化的值自动设置长度

	fmt.Println("arr[2:6]",arr[2:6])
	fmt.Println("arr[:6]",arr[:6])
	fmt.Println("arr[2:]",arr[2:])
	fmt.Println("arr[:]",arr[:])
	s1:=arr[2:] //切片得到的新变量, 是引用类型
	s2:=arr[:]

	fmt.Println("before update slice s1", s1)
	fmt.Println("before  update slice s2", s2)

	updataSlice(s1)
	updataSlice(s2)
	fmt.Println("after update slice s1", s1)
	fmt.Println("after update slice s2", s2)
	fmt.Println("after update slice arr", arr)

	s3:=s1[2:]
	fmt.Println("resize s1, s3 = ", s3)

	fmt.Println("extending slice") //
	arr=[...]int{0,1,2,3,4,5,6,7}
	s4 :=arr[2:6]
	s5 :=s4[3:5]
	fmt.Printf("s4=%v, len(s4) =%d, cap(s4)=%d\n", s4,len(s4), cap(s4))
	fmt.Printf("s5=%v, len(s5) =%d, cap(s5)=%d\n", s5,len(s5), cap(s5))

	s6:=append(s4,10)
	s7:=append(s6,11)//超出了原数组的索引了, 会copy一份原数组, 由于开辟了新空间, 必须要有新的切片来接收
	s8:=append(s7,12)
	fmt.Println("s6, s7 ,s8", s6, s7, s8)
	fmt.Println("arr", arr)
}
