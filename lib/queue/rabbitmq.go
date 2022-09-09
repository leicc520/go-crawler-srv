package queue

import (
	"context"
	"encoding/json"

	"github.com/leicc520/go-orm/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqSt struct {
	Url string `json:"url" yaml:"url"`
	Durable bool `json:"durable" yaml:"durable"`
	AutoAck bool `json:"auto_ack" yaml:"auto_ack"`
	AutoDelete bool `json:"auto_delete" yaml:"auto_delete"`
	Queue string `json:"queue" yaml:"queue"`
	conn *amqp.Connection
	ch   *amqp.Channel
	q     amqp.Queue
}

//往队列发送消息 允许直接传递消息对象指针
func (s *RabbitMqSt) Publish(data interface{}) error {
	var ok = false
	var pMsg *amqp.Publishing = nil
	if pMsg, ok = data.(*amqp.Publishing); !ok {
		//字符串直接转字节数组即可
		var body []byte = nil
		if str, ok := data.(string); ok {
			body = []byte(str)
		} else {
			body, _ = json.Marshal(data)
		}
		pMsg = &amqp.Publishing{ContentType: "text/plain", Body: body}
	}
	ctx := context.Background()
	err := s.ch.PublishWithContext(ctx, "", s.Queue, false, false, *pMsg)
	if err != nil {//失败的情况处理逻辑
		log.Write(log.ERROR, "Failed to publish a message", err)
	}
	return err
}

//克隆一个对象处理逻辑
func (s *RabbitMqSt) Clone(queue string) *RabbitMqSt{
	c := &RabbitMqSt{Url: s.Url, Durable: s.Durable,
		AutoAck: s.AutoAck, AutoDelete: s.AutoDelete, Queue:s.Queue+queue}
	return c
}

//初始化队列的链接处理逻辑
func (s *RabbitMqSt) Init() error {
	var err error
	s.conn, err = amqp.Dial(s.Url)
	if err != nil {
		log.Write(log.ERROR, "Failed connect ", err)
		return err
	}
	s.ch, err = s.conn.Channel()
	if err != nil {
		log.Write(log.ERROR, "Failed to open a channel", err)
		return err
	}
	s.q, err = s.ch.QueueDeclare(
		s.Queue, // name
		s.Durable,   // durable
		s.AutoDelete, // delete when unused
		false,   // exclusive
		false,  // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Write(log.ERROR, "Failed to declare a queue", err)
		return err
	}
	return nil
}

//关闭释放处理逻辑
func (s *RabbitMqSt) Close()  {
	if s.ch != nil {
		s.ch.Close()
	}
	if s.conn != nil {
		s.Close()
	}
}

//消费队列数据资料信息
func (s *RabbitMqSt) Consumer(handle QueueCB) error {
	msgChan, err := s.ch.Consume(
		s.q.Name, // queue
		"",     // consumer
		s.AutoAck,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Write(log.ERROR, "Failed to register a consumer", err)
		return err
	}
	for {
		dMsg, isClose := <-msgChan
		if !isClose {
			log.Write(log.ERROR, s.Queue, "queue closed")
			break
		}
		err = handle(dMsg.Body) //业务只需关注数据即可
		if !s.AutoAck { //如果是非自动确认的话 需要手动确认
			if err == nil {
				s.ch.Ack(dMsg.DeliveryTag, false)
			} else {
				s.ch.Reject(dMsg.DeliveryTag, true)
			}
		}
		log.Write(log.INFO, s.Queue, " Received a message: ", dMsg.Body, err)
	}
	return err
}

//消费队列数据资料信息
func (s *RabbitMqSt) AsyncConsumer(conCurrency int, handle QueueCB) error {
	msgChan, err := s.ch.Consume(
		s.q.Name, // queue
		"",     // consumer
		s.AutoAck,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Write(log.ERROR, "Failed to register a consumer", err)
		return err
	}
	//设置缓冲区 以便于控制开启协程的数量，限制并发
	goConCurrencyChan := make(chan byte, conCurrency)
	for {
		dMsg, isClose := <-msgChan
		if !isClose {
			log.Write(log.ERROR, s.Queue, "queue closed")
			break
		}
		goConCurrencyChan <- 1
		go func(dlMsg amqp.Delivery) {
			defer func() { //结束释放并发位置
				<-goConCurrencyChan
			}()
			err = handle(dlMsg.Body) //业务只需关注数据即可
			if !s.AutoAck { //如果是非自动确认的话 需要手动确认
				if err == nil {
					s.ch.Ack(dlMsg.DeliveryTag, false)
				} else {
					s.ch.Reject(dlMsg.DeliveryTag, true)
				}
			}
			log.Write(log.INFO, s.Queue, " Received a message: ", dMsg.Body, err)
		}(dMsg)
	}
	return err
}
