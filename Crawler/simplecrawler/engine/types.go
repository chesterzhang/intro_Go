package engine

type Request struct {
	Url string //城市 Url
	ParserFunc func([] byte) ParseResult //一个函数, 返回对html文本的解析结果

}

type ParseResult struct {
	Requests []Request
	Items []interface{}  //城市名
}

// 一个伪 ParseFunc, 后面会用来
func NilParser([] byte) ParseResult{
	return ParseResult{}
}


