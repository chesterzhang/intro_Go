package persist

import (
	"Crawler/concurrentcrawler/types"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"

	//"log"
)

// 将 item 存起来

// 从一个 channel 里面一直去读
func ItemSaver(index string) (chan types.Item, error)   {
	out:=make(chan types.Item)

	client,err:=elastic.NewClient(elastic.SetSniff(false))  // 因为跑在本地dockers上面, 必须设置为 false
	if err!=nil{
		return  nil,err
	}

	go func() {
		itemCount :=0
		for{
			item := <- out
			//log.Printf("Item saver: got item %d, %s : ", itemCount, item )

			//存入 elasticserch
			_,err:=save(index, client,item)
			if err != nil {
				log.Printf("Item Saver: erro" +"saving item %v: %v", item,err)
			}

			fmt.Println(item)


			itemCount++
		}
	}()
	return  out, nil
}

func save( index string,client *elastic.Client,item interface{} ) (id string, err error ){

	// Index(数据库名), Type(表名)
	resp,err:=client.Index().
		Index(index).
		Type("zhenai").
		BodyJson(item).Do(context.Background())

	if err!=nil {
		return "", err
	}

	// %+v 如果打印结构体, 会将结构体的字段名打出来
	fmt.Printf("Saved item: %+v \n", item)

	return  resp.Id, nil
}
