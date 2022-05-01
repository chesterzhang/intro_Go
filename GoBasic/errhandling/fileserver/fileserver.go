package main

import (
	"filelisting"
	"log"

	"os"

	"net/http"

)

// 定义一个 函数类型, 名叫 appHandler, 返回 error 类型
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 将 filelisting.HandleFileList 传进去作为 handler, 并且返回一个 函数
func errWrapper(handler appHandler)  func(writer http.ResponseWriter, request *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {

		//自己recover, 对下面 handler 传回来的err 进行处理
		defer func() {
			r:=recover()
			if r!=nil{// 因为 recover() 返回一个 interface{} 也就是 任何类型, 所以要判断一下
				log.Println("pannic :",r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

		}()

		err:= handler(writer, request) // 本质上是 filelisting.HandleFileList(writer, request)

		if err!=nil{
		 	log.Println(err)

		  	userErr, ok:=err.(userError)// 判断一下, 是不是我们自定义的 userError 类型
			if ok {
				http.Error(writer,userErr.Message(),http.StatusBadRequest)
				return
			}

			code :=http.StatusOK
			switch   {
			case os.IsNotExist(err):
				code=http.StatusNotFound
			case os.IsPermission(err):
				code=http.StatusForbidden
			default:
				code=http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

//自定义 error, 一个接口
type userError interface {
	error  // 接口里面的接口实现, 是一个error
	Message()  string // 接口里面的一个函数
}

func main() {

	// 为 指定的 url 注册一个 handler, 负责接受请求, 第二个参数为一个 handler func(ResponseWriter, *Request)
	http.HandleFunc("/",errWrapper(filelisting.HandleFileList))


	err:=http.ListenAndServe(":8888",nil)
	if err!=nil {
		panic(err)
	}
	
}
