package types

// engine 的功能就是 接收返回的 ParseResult, 和 分发 request

type Request struct {
	Url string //城市 Url
	ParserFunc func([] byte) ParseResult //一个函数, 返回对html文本的解析结果

}

type ParseResult struct {
	Requests []Request
	Items [] Item  //城市名 或者 用户名
}

// 需要存入 ElasticSearch 的 一个个对象
type Item struct {
	ItemName string // 城市名, 用户名
	ItemType string //city or user, 表示 item 类型, 是城市 还是用户
}


// 一个伪 ParseFunc, 当不想进一步向下解析时用到
func NilParser([] byte) ParseResult{
	return ParseResult{}
}

// scheduler , 用来具体实现如何接收 result, 分发 request
//一个 interface 收集 request, 派发request 给worker
type  Scheduler interface {

	// 向 worker 传送 request
	// simple 并发版本, 向所有 worker 公用的 channel 也就是 workerInChan 写入 一个 Request
	Submit(Request)

	// 每一个worker, 问scheduler要channel
	// simple 并发版本 return 一个所有worker 公用的 channel(这个channel 写在结构体的成员属性, 由调用 Run 方法时创建好的)
	// queue 并发版本 每个worker 调用一次 这个方法 就 创建一个新的 channel 然后 再 return 回去
	CreateWorkerChan() chan Request

	//ConfigurateWorkerChan(chan Request)//将worker要用到获取request的channel给scheduler,  scheduler 往里面写request

	// simple 并发版本不实现这个方法
	// queue 并发版本 中 每一个 worker 会将自己的 channel 交给 scheduler 排队, 等待scheduler 分发 request
	WorkerReady(chan Request)

	// 初始化 scheduler
	// simple 并发版本就 make 一个所有worker公用的 channel, 用于给 worker 传入 Request
	// queue 并发版本 则会初始化 request 队列, chan request 队列(每一个worker的)
	// 不停地 收 request 入队, chan request 入队, 如果 两个队列都不为空, 就取出 request 队列队首, 取出 chan request 队列队首
	// 并且将 request 队首送入 chan request 队首(相当于给worker分发 request)
	Run()
}



