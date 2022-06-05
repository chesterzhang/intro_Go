package parser

import (
	"io/ioutil"
	"testing"
)



func TestParseCityList(t *testing.T) {
	contents, err:=ioutil.ReadFile("citylist_test_data.html")
	if err!=nil {
		panic(err)
	}

	result:=ParseCityList(contents)
	const resultSize  int = 470
	urlExpected :=[]string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	cityExpected :=[] string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}

	if len(result.Items)!=resultSize {
		t.Errorf("Result.Items size should be %d, but got %d \n",resultSize,len(result.Items))
	}

	if len(result.Requests)!=resultSize {
		t.Errorf("Result.Requests size should be %d, but got %d \n",resultSize,len(result.Requests))
	}

	for i,url := range  urlExpected{
		if url!=result.Requests[i].Url {
			t.Errorf("Url shoudl be %s, but got %s \n",url,result.Requests[i].Url)
		}
	}

	for i,city := range  cityExpected{
		if city!=result.Items[i] {
			t.Errorf("City shoudl be %s, but got %s \n",city,result.Items[i])
		}
	}

}
