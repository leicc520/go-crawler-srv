package queue

import (
	"context"
	"encoding/json"

	"github.com/leicc520/go-orm/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqSt struct {
	Url string `json:"url" yaml:"url"`
	Queue string `json:"queue" yaml:"queue"`
	conn *amqp.Connection
	ch   *amqp.Channel
	q     amqp.Queue
}

//往队列发送消息
func (s *RabbitMqSt) Publish(data interface{}) error {
	body, _ := json.Marshal(data)
	ctx  := context.Background()
	err  := s.ch.PublishWithContext(ctx, "", s.Queue, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: body})
	if err != nil {//失败的情况处理逻辑
		log.Write(log.ERROR, "Failed to publish a message", err)
	}
	return err
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
		false,   // durable
		false, // delete when unused
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
	if s.conn != nil {
		s.Close()
	}
	if s.ch != nil {
		s.ch.Close()
	}
}

//消费队列数据资料信息
func (s *RabbitMqSt) Consumer(handle QueueCB) error {
	msgChan, err := s.ch.Consume(
		s.q.Name, // queue
		"",     // consumer
		true,   // auto-ack
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
		err = handle(dMsg.Body)
		log.Write(log.INFO, s.Queue, " Received a message: ", string(dMsg.Body), err)
	}
	return err
}
