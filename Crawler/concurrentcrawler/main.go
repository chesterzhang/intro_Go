package main

import (
	"Crawler/concurrentcrawler/engine"
	"Crawler/concurrentcrawler/persist"

	//"Crawler/concurrentcrawler/persist"
	"Crawler/concurrentcrawler/scheduler"
	"Crawler/concurrentcrawler/types"
	"Crawler/concurrentcrawler/zhenai/parser"
)

func main() {
	itemSaveChan, err := persist.ItemSaver("dating_profile")

	if err!=nil {
		panic(err)
	}

	 //e:=engine.ConcurentEngine{Scheduler:&scheduler.SimpleScheduler{},WorkerCount:5 ,ItemSaveChan : itemSaveChan}
	e:=engine.QueueConcurentEngine{ Scheduler:&scheduler.QueueScheduler{},WorkerCount:5, ItemSaveChan: itemSaveChan }
	e.Run( types.Request{Url: "http://www.zhenai.com/zhenghun",ParserFunc:parser.ParseCityList} )

	//log.Print("ok")
}


