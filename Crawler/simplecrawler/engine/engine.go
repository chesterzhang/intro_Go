package engine

import (
	"Crawler/simplecrawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r:=range seeds{
		requests=append(requests,r)
	}

	for len(requests) >0 {
		r:=requests[0]//将第一个 request 取出来

		requests=requests[1:] // 相当于队列.pop

		body, err :=fetcher.Fetch(r.Url)// 从 URL 中提取 HTML 文本
		if err!=nil {
			log.Printf("fetcher : error "+
				"fetching url %s: %v", r.Url, err)
			continue
		}
		log.Printf("Feteching URL: %s \n", r.Url)  //OKHERE


		parseResult:= r.ParserFunc(body) // 用request对应的 parseFunc 解析出result
		requests=append(requests, parseResult.Requests...)

		for _,item := range parseResult.Items{
			log.Printf("Got item : %s \n",item)
		}

	}
}


