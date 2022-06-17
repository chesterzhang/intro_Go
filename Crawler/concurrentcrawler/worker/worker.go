package worker

import (
	"Crawler/concurrentcrawler/fetcher"
	"Crawler/concurrentcrawler/types"
	"log"
)

func Worker(r types.Request) (types.ParseResult, error) {
		body, err :=fetcher.Fetch(r.Url)// 通过 URL 获取网页的 HTML 文本
		if err!=nil {
			log.Printf("fetcher : error "+
				"fetching url %s: %v", r.Url, err)
			return types.ParseResult{},err
		}
		//log.Printf("Feteching URL: %s \n", r.Url)

		// 用request对应的 parseFunc(可能是 ParseCityList, ParseCity) 解析 HTML 返回 result
		parseResult:= r.ParserFunc(body)

		return  parseResult,nil
}

