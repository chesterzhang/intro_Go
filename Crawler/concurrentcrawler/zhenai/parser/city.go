package parser

import (
	"Crawler/concurrentcrawler/types"
	"regexp"
)

//在每一个城市的URL中将所有用户提取出来

//<a href="http://album.zhenai.com/u/1958903678" target="_blank">小扬姐</a>
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)[^>]*>([^>]+)</a>`


func ParseCity(contents []byte) types.ParseResult {
	re:=regexp.MustCompile(cityRe )

	matches:=re.FindAllSubmatch(contents,-1) // [][][]byte 所有用户URL 和 用户名称
	//fmt.Println(matches)
	result:= types.ParseResult{} // 最后return出去的ParseResult

	for _,m:=range matches {// m[2]用户名, m[1]用户URL
		//result.Items = append(result.Items, string(m[2])) //将用户名append进去
		result.Items = append(result.Items,   types.Item{ItemName: string(m[2]), ItemType:"city"}) //将用户名append进去
		result.Requests=append(result.Requests, types.Request{
			Url:        string(m[1]),    // []byte 转换成 string
			ParserFunc: types.NilParser, //这里的parserFunc 应该是从用户URL里面提取用户详细信息的 parser
		})
	}

	//最后的 result 的 items为所有用户名的slice, 和request的slice{用户1, 用户URL解析用户详细信息的parseFunc(暂时为空),...}

	//fmt.Println("len(matches):",len(matches))
	return result
}
