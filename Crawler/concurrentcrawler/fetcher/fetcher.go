package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(40*time.Millisecond)//  每隔40ms爬一次, 以防止被反爬

// 给定一个URL, 返回其页面的 HTML 文本
func Fetch(url string) ([] byte, error)  {
	<- rateLimiter // 每隔40ms爬一次, 以防止被反爬
	resp,err:=http.Get(url)
	if err!=nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK{

		return nil, fmt.Errorf("wrong status code : %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)

}
