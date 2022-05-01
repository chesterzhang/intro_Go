package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (e userError) Error() string  {
	return e.Message()
}

func (e userError) Message() string  {
	return  string(e)
}

const prefix  ="/list/"
func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix)!=0{
		return  userError("path must start with "+prefix)
	}
	path :=request.URL.Path[len("/list/"):] //  /list/file.txt 提取出 file.txt
	file, err:=os.Open(path)
	if err!=nil {//如果打不开文件
		//http.Error(writer,err.Error(),http.StatusInternalServerError)
		return err
	}

	defer  file.Close()

	all, err:=ioutil.ReadAll(file)
	if err !=nil {
		return  err
	}
	writer.Write(all)
	return nil
}
