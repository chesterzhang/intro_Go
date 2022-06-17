package engine

import (
	"Crawler/concurrentcrawler/types"
	"Crawler/concurrentcrawler/worker"
	"fmt"
)

type QueueConcurentEngine struct {
	Scheduler   types.Scheduler //sheduler(收集 request, 派发request 给worker)
	WorkerCount int             //worker 数量
	ItemSaveChan  chan types.Item // 用于传入 item,  然后将这个 item 写入 elasticSearch
}

////一个 interface 收集 request, 派发request 给worker
//type QueueScheduler interface {
//	Submit(Request)
//	WorkerChan() chan Request // 有一个worker, 问scheduler要channel
//	//ConfigurateWorkerChan(chan Request)//将worker要用到获取request的channel给scheduler,  scheduler 往里面写request
//	WorkerReady(chan Request)
//	Run()
//}

func (e *QueueConcurentEngine ) Run(seeds ...types.Request)  {

	// 创建多个goroutine, 每一个goroutine 里面都要运行一个worker(fetcher+parser)
	// 每一个 goroutine 不停地从 chan Request 中读 request, 由 worker(fetcher+parser 处理好以后) 向 chanParseResult 中写 ParseResult
	// goroutine 获得 channel 的权限由
	// workerInChan:=make(chan Request) //从 schedule 获取 request的channel
	workerOutChan:=make(chan types.ParseResult) //多个 worker 输出共用的 channel

	//fmt.Println("ok here") //ok here

	e.Scheduler.Run()//

	//fmt.Println("ok here") // BUG ABOVE
	for i:=0;i<e.WorkerCount ;i++  {
		createQueueWorker(e.Scheduler.CreateWorkerChan(), workerOutChan, e.Scheduler)//channel 是引用类型,无需传递指针
	}


	//将起始 request 传给 scheduler
	for _, r:= range seeds{
		e.Scheduler.Submit(r)//需要先ConfigurateWorkerChan,指定一个chan, 同时这个chan 由  才能submit
	}

	for {
		result:= <-workerOutChan
		// 不应该在这里打印输出, 因为 engine 需要尽可能快的处理 request
		//for _,item:=range result.Items{
		//	fmt.Println(item)
		//}

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

//创建一个 goroutine, 不停地从 chan Request 中读 request, 由 worker(fetcher+parser 处理好以后) 向 chanParseResult 中写 ParseResult
func createQueueWorker(in chan  types.Request, out chan types.ParseResult,s types.Scheduler)  {

	go func() {
		for{
			// 每一个 worker 将自己用于获取 request 的 chanel 交给 scheduler, 等待 scheduler 分发 request
			s.WorkerReady(in)
			request:= <- in// 等待 scheduler 分发 request
			//result,err:=worker.Worker.Do(request)
			result,err:=worker.Worker(request)

			if err!=nil {
				continue
			}
			for _, item:=range result.Items{
				fmt.Printf("%s \n",item)
			}
			out<- result
		}
	}()
}