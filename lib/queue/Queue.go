package queue

type QueueCB func([]byte) error

//定义任务队列的处理逻辑
type IFQueue interface {
	Close()
	Init() error
	Publish(data interface{}) error
	Consumer(handle QueueCB) error
}
