package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type MyRetriever struct {
	UserAgent string
	TimeOut time.Duration
	CourseName string
}

// 相当于 MyRetriver 实现 Get 方法
func (r *MyRetriever) Get(url string) string  {
	resp, err := http.Get(url)
	if err!=nil {
		panic(err)
	}

	result, err:=httputil.DumpResponse(resp,true)
	resp.Body.Close()
	if err!=nil {
		panic(err)
	}

	return  string(result)
}

// 定义一个接口, 接口里面有 Get 方法
type Retriever interface {
	Get(url string) string
}

// 定义一个接口, 接口里面有 Post 方法
type Poster interface{
	Post(url string, form map[string] string) string
}

// 定义一个接口, 接口里面既包含Retriever 接口,也包含Poster接口
type RetrieverPoster interface {
	Retriever
	Poster
}

//MyRetriever 也实现了 Post 方法, 相当于实现了 Poster, Retriver 两个接口
func (r *MyRetriever)  Post(url string, form map[string] string) string{
	r.CourseName=form["CourseName"]
	return "ok"
}

// 实现一个 go 语言常用系统接口, String() 方法相当于 Java 的toString 方法
func (r *MyRetriever)  String() string{
	return fmt.Sprintf("MyRetriver:{UserAgent=%s,TimeOut=%s,CourseName=%s}",r.UserAgent,r.TimeOut,r.CourseName)
}

func main() {
	var r Retriever
	r=&MyRetriever{UserAgent:"Mozilla/5.0", TimeOut: time.Minute}
	fmt.Printf("%T %v\n",r ,r) //输出类型, 值

	// type assertion 查看 接口变量 类型
	mr,ok:=r.(*MyRetriever)
	if ok {
		fmt.Println(mr.UserAgent,mr.TimeOut)
	}else {
		fmt.Println("Not a my receiver")
	}

	// type switch 查看 接口变量 类型
	switch tp:=r.(type) {
	case *MyRetriever:
		fmt.Println(tp.UserAgent,tp.TimeOut)
	default:
		fmt.Println("Not a my receiver")
	}

	fmt.Println(r) //只要实现了 String()方法,就可以直接打印

	//fmt.Println(r.Get("http://wwwimooc.com"))

	fmt.Println("=======接口组合=========")
	PostMap :=map[string] string{
		"CourseName": "Golang",
	}
	var r2 RetrieverPoster
	r2 = &MyRetriever{UserAgent:"Mozilla/5.0", TimeOut: time.Minute}
	fmt.Println(r2.Post("http://www.imooc.com",PostMap))
	fmt.Printf("%T %v\n",r2 ,r2) //输出类型, 值

}
