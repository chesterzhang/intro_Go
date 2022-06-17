package engine

import (
	"Crawler/concurrentcrawler/types"
	"Crawler/concurrentcrawler/worker"
)

type ConcurentEngine struct {
	 Scheduler   types.Scheduler //sheduler(收集 request, 派发request 给worker)
	 WorkerCount int             //worker 数量
	 ItemSaveChan  chan types.Item // 用于传入 item,  然后将这个 item 写入 elasticSearch
}

//一个 interface 收集 request, 派发request 给worker
//type Scheduler interface {
//	Submit(Request)
//	ConfigurateWorkerChan(chan Request)//将worker要用到获取request的channel给scheduler,  scheduler 往里面写request
//}

func (e *ConcurentEngine ) Run(seeds ...types.Request)  {

	e.Scheduler.Run()// 在 scheduler 内创建一个 request channel 用于向多个 worker 发送 request
	workerOutChan:=make(chan types.ParseResult) //多个 worker 输出 result 的共用的 channel

	//创建多个goroutine, 每一个goroutine 里面都要运行一个worker(fetcher+parser)
	//每一个 goroutine 不停地从 chan Request 中读 request, 由 worker(fetcher+parser 处理好以后) 向 workerOutChan 中写 ParseResult

	for i:=0;i<e.WorkerCount ;i++  {
		createWorker(e.Scheduler.CreateWorkerChan()  ,workerOutChan)//channel 是引用类型,无需传递指针
	}

	//将起始 request{URL, ParseFunc}传给 scheduler
	for _, r:= range seeds{
		e.Scheduler.Submit(r)//需要先ConfigurateWorkerChan,指定一个chan, 同时这个chan 由  才能submit
	}

	for {

		// 不停地从  workerOutChan 中 读取 result
		result:= <-workerOutChan

		//不应该在这里打印输出(这里仅仅为了调试),  因为 engine 需要尽可能快的处理 request
		//for _,item:=range result.Items{
		//	fmt.Println(item)
		//}

		//
		for _,item:=range result.Items{

			 go func( item types.Item) {
			 	e.ItemSaveChan<- item // 将 item 不停地送入一个 channel, 写入 elasticSearch
			 }(item)

		}

		for _, request :=range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

//创建一个 goroutine, 不停地从 Scheduler 的 workInChan 中读 request, 由 worker(fetcher+parser 处理好以后) 向 chanParseResult 中写 ParseResult
func createWorker(in chan types.Request, out chan types.ParseResult)  {
	go func() {
		for{
			request:= <- in // 从 Scheduler 的 workInChan 中读 request
			result,err:=worker.Worker( request) // 解析 request, 返回 result(request 里面URL网页的 item, 和进一步的 ParseFunc)
			if err!=nil {
				continue
			}
			//for _, item:=range result.Items{
			//	fmt.Printf("%s \n",item)
			//}
			out<- result// 将 result 写入 engine 的 workerOutChan
		}
	}()
}

