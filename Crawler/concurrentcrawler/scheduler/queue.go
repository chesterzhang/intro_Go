package scheduler

import (
	"Crawler/concurrentcrawler/types"
)

type QueueScheduler struct {
	requestChan chan types.Request // scheduler 收到的 request

	// 负责接收 每一个 worker 都独有的 用于接收 request的channel
	// 然后将这些 channel, 存进一个队列, 获取到了这些 worker的 channel, 也就相当于获取了worker
	workerChan chan chan types.Request
}

func (s *QueueScheduler) CreateWorkerChan() chan types.Request {
	// 每一个 worker 都有一个 chan Request, 每调用一次这个方法, 都新创建一个channel return回去
	return  make(chan types.Request)
}

// 将 request 送入 一个 worker 的
func (s *QueueScheduler) Submit(r types.Request) {
	s.requestChan <- r
}

// 每一个 worker 会将自己的 channel 交给 scheduler 排队, 等待scheduler 分发 request
func (s *QueueScheduler) WorkerReady(w chan types.Request) {
	s.workerChan <- w
}


// 初始化 scheduler
// 初始化 request 队列, chan request 队列(每一个worker的chan)
// 不停地 收 request 入队, chan request 入队, 如果 两个队列都不为空, 就取出 request 队列队首, 取出 chan request 队列队首
// 并且将 request 队首送入 chan request 队首(相当于给worker分发 request)
func (s *QueueScheduler) Run() {
	s.requestChan=make(chan types.Request)
	s.workerChan =make(chan chan types.Request)

	go func() {
		var requestQ [] types.Request      // 存储 requeset的队列
		var workerQ  [] chan types.Request // 间接的 worker 队列, 拿到了每一个 worker 用来读 request的 chan, 相当于获取到了worker
		for {
			var activeRequest types.Request
			var activeWorker chan types.Request
			if len(requestQ)>0 && len(workerQ )>0 {
				activeRequest = requestQ[0]// 将队首的 request 取出来
				activeWorker = workerQ[0] // 将队首的 worker 的 channel 取出来
			}
			select {
			case r:= <- s.requestChan:// 如果有 request 进来了
				requestQ=append(requestQ, r)
			case w:= <- s.workerChan: // 如果有 worker 告诉 scheduler 准备好了(调用了 WorkerReady 函数)
				workerQ = append(workerQ,w)// 将 worker 的 chan rereques 存起来
			case activeWorker <- activeRequest: // 将 request 分发给了 worker 获取 request 的channel
				workerQ=workerQ[1:]
				requestQ=requestQ[1:]
			}
		}
	}()


}

