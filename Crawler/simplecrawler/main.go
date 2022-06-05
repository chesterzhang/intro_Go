package main

import (
	"Crawler/simplecrawler/engine"
	"Crawler/simplecrawler/zhenai/parser"

)

func main() {
	 engine.Run(engine.Request{Url:"http://www.zhenai.com/zhenghun",ParserFunc:parser.ParseCityList})

}


