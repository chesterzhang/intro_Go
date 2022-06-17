package scheduler

import (
	"Crawler/concurrentcrawler/types"
)

// 一个简易版本 scheduler
// 当有 request 进来时, 将 request 送入 一个 channel, 所有worker 都从这个channel 获取 request
type SimpleScheduler struct {
	workerInChan chan types.Request
}

// 创建每一个 worker
func (s *SimpleScheduler) CreateWorkerChan() chan types.Request {
	return  s.workerInChan
}

//simple 并发版 scheduler 不需要实现这个接口
func (s *SimpleScheduler) WorkerReady(chan types.Request) {
	//panic("implement me")
}

// 初始化 scheduler, 创建好公用的 channel
func (s *SimpleScheduler) Run() {
	s.workerInChan =make(chan types.Request)
}

//将 收到的 request 送到 workerInchan 中去
func (s *SimpleScheduler) Submit(r types.Request) {
	// 这里必须用一个 go func, 否则会产生死锁!!!
	// 如果不用 go func, 当所有 worker 都没空接收 request时, 就会卡在这里
	// 对于无缓冲区的 channel 来说, 读写必须同步
	// 而worker 想要腾出手来接收 worker 就必须进入engine里的下一个 循环把 result 取出来
	// request 等待 空闲的worker接收, 忙碌的 worker 等待下一个循环将result送出, 循环等待由此产生
	//s.workerInChan<- r
	go func() {
		s.workerInChan<- r
	}()
}

////指定向哪一个channel写入 request
//func (s *SimpleScheduler) ConfigurateWorkerChan(c chan engine.Request) {
//	s.workerInChan=c
//}
