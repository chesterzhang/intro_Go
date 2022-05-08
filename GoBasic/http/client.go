package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err:=http.NewRequest(http.MethodGet,"http://www.imooc.com",nil )
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	//第一个参数 request, 第二个参数是重定向 request的路径
	client:=http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		fmt.Println("Redirect:", req)
		return  nil//正常重定向
	}}

	resp,err :=client.Do(request)
	//resp,err :=http.DefaultClient.Do(request)
	//resp,err :=http.Get("http://www.imooc.com")
	if err !=nil{
		panic(err)
	}
	defer resp.Body.Close()//最后一定要关闭

	s,err:=httputil.DumpResponse(resp,true)

	if err!=nil{
		panic(err)
	}

	fmt.Printf("%s \n",s)

}

