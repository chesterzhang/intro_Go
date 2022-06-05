package parser

import (
	"Crawler/simplecrawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 将城市列表(也就是所有城市)用正则表达式提取出来
//<a href="http://www.zhenai.com/zhenghun/baoshan1" data-v-1573aa7c>保山</a>
// contents 也就是 html 文本, 从文本中提取出 城市URL 和城市名称
func ParseCityList(contents []byte) engine.ParseResult{
	re:=regexp.MustCompile(cityListRe)
	//re:=regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/*`) //OK here

	matches:=re.FindAllSubmatch(contents,-1) // [][][]byte 所有城市URL 和城市名
	//fmt.Println(matches)
	result:=engine.ParseResult{}// 最后return出去的ParseResult

	for _,m:=range matches {// m[2]城市名, m[1]城市URL
		result.Items = append(result.Items, string(m[2])) //将城市名称append进去
		result.Requests=append(result.Requests,engine.Request{
			Url:string(m[1]),// []byte 转换成 string
			//ParserFunc: ParseCity, //这里的parserFunc 应该是从城市URL里面提取用户URL的 parser
			ParserFunc: engine.NilParser,
		})
	}

	//最后的 result 的 items为所有城市的slice, 和request的slice{城市1, 城市解析用户的parseFunc(暂时为空),...}

	//fmt.Println("len(matches):",len(matches))
	return result
}



