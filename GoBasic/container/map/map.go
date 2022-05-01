package main

import "fmt"

func main() {
	// 一个 hashMap
	m :=map[string] string{
		"site": "google.com",
		"location" : "US",
	}

	m2 :=make(map[string] int) //一个空map
	fmt.Println(m,m2)

	fmt.Println("Traversing map")
	for k,v := range m{
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	url,ok:= m["site"]// 若 key 存在, 则返回ok为 true, 否则为 false
	fmt.Println("url", url,ok)

	url,ok= m["s1te"]
	fmt.Println("url", url,ok) //map 中不存在key 会返回一个空

}
