package persist

import (
	"Crawler/concurrentcrawler/types"
	"encoding/json"

	//"Crawler/concurrentcrawler/types"
	"context"
	"github.com/olivere/elastic/v7"
	"testing"
)

// 单例测试能否 向 elasticSearch 里面存入一个对象
func TestSaver(t *testing.T) {
	item := types.Item{
		ItemName:"重庆",
		ItemType: "city",
	}

	// 尝试将 item 存入 elasticSearch
	//id, err:=save(item)
	//
	//if err!=nil{
	//	panic(err)
	//}
	//
	// 尝试 根据 id 去 elasticSearch 读一个 item
	client,err:=elastic.NewClient(elastic.SetSniff(false))  // 因为跑在本地dockers上面, 必须设置为 false
	if err!=nil{
		panic(err)
	}

	resp,err:=client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id("sjF_GoEBJWvoM86p2vtn").Do(context.Background())

	if err!=nil {
		panic(err)
	}

	//t.Logf("%s",resp.Source)

	var actual types.Item
	err=json.Unmarshal([]byte (resp.Source), &actual)

	if err!=nil {
		panic(err)
	}

	if actual != item {
		t.Errorf("Got %v ; EXPECTED %v", actual, item)
	}

}