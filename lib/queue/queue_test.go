package queue

import (
	"github.com/leicc520/go-orm"
	"github.com/leicc520/go-orm/log"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := IFQueue(&RabbitMqSt{Url: "amqp://guest:guest@10.100.72.102:5672/", Queue: "demo"})

	err := q.Init()
	if err != nil {
		return
	}
	defer q.Close()
	cb := func([]byte) error { return nil}
	go q.Consumer(cb)

	sp := func(i int, max int) {
		for i < max {
			q.Publish(orm.SqlMap{"data":i})
			log.Write(log.INFO, i)
			i++
			time.Sleep(time.Millisecond*50)
		}
	}

	go sp(1, 1000000)
	go sp(30, 50)
	go sp(1, 20)
	go sp(30, 50)
	go sp(1, 20)
	go sp(30, 50)
	go sp(1, 20)
	go sp(30, 50)
	go sp(1, 20)
	go sp(30, 50)
	go sp(1, 20)
	go sp(30, 50)

	sp(60, 100)

	c := make(chan int)
	<-c
}
