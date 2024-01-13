package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"new-fix/types"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	err                     error
	rabbitMqConn            *amqp.Connection
	rabbitMqChannel         *amqp.Channel
	queueDeclare            amqp.Queue
	productionQueueDeclares = make(map[string]amqp.Queue)
	consumeQueueDeclares    = make(map[string]struct {
		queue    amqp.Queue
		messages <-chan amqp.Delivery
	})
	ctx      context.Context
	cancelFn context.CancelFunc
)

// RabbitMqSendMessage 发送rabbit-mq消息
func RabbitMqSendMessage(queueName, message string) {
	publishContext := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}
	if err = rabbitMqChannel.PublishWithContext(ctx,
		"",                                      // exchange
		productionQueueDeclares[queueName].Name, // routing key
		false,                                   // mandatory
		false,                                   // immediate
		publishContext); err != nil {
		log.Printf("[rabbit-mq-error] [发起消息失败] %v\n", err)
	} else {
		log.Printf("[rabbit-mq-debug] [发起消息成功] %s %v\n", productionQueueDeclares[queueName].Name, publishContext)
	}
}

// RabbitMqHandler rabbit-mq处理器
func RabbitMqHandler(
	rabbitMqUsername,
	rabbitMqPassword,
	rabbitMqAddr,
	rabbitMqVhost string,
	productionQueueNames,
	consumeQueueNames []string,
) {
	rabbitMqConn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", rabbitMqUsername, rabbitMqPassword, rabbitMqAddr, rabbitMqVhost))
	if err != nil {
		log.Fatalf("[rabbit-mq-error] [创建连接失败] %v\n", err)
	}
	log.Printf("[rabbit-mq-debug] [创建连接成功]\n")
	defer func(conn *amqp.Connection) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("[rabbit-mq-error] [关闭连接失败] %v\n", err)
		}
	}(rabbitMqConn)

	rabbitMqChannel, err = rabbitMqConn.Channel()
	if err != nil {
		log.Printf("[rabbit-mq-error] [创建信道失败] %v\n", err)
	}
	log.Printf("[rabbit-mq-debug] [创建信道成功]")
	defer func(ch *amqp.Channel) {
		err = ch.Close()
		if err != nil {
			log.Printf("[rabbit-mq-error] [关闭信道失败] %v\n", err)
		}
	}(rabbitMqChannel)

	// 链接队列（生产）
	for _, productionQueueName := range productionQueueNames {
		queueDeclare, err := rabbitMqChannel.QueueDeclare(
			productionQueueName, // name
			false,               // durable
			false,               // delete when unused
			false,               // exclusive
			false,               // no-wait
			nil,                 // arguments
		)
		if err != nil {
			log.Printf("[rabbit-mq-error] [链接队列失败] %v\n", err)
		}
		log.Printf("[rabbit-mq-debug] [链接队列成功] %s", queueDeclare.Name)
		productionQueueDeclares[queueDeclare.Name] = queueDeclare // 保存队列
	}

	// 链接队列（消费）
	for _, consumeQueueName := range consumeQueueNames {
		queueDeclare, err = rabbitMqChannel.QueueDeclare(
			consumeQueueName, // name
			false,            // durable
			false,            // delete when unused
			false,            // exclusive
			false,            // no-wait
			nil,              // arguments
		)
		if err != nil {
			log.Printf("[rabbit-mq-error] [链接队列失败] %v\n", err)
		}
		log.Printf("[rabbit-mq-debug] [链接队列成功] %s", queueDeclare.Name)

		// 创建监听
		receiveMessages, err := rabbitMqChannel.Consume(
			queueDeclare.Name, // queue
			"",                // consumer
			true,              // auto-ack
			false,             // exclusive
			false,             // no-local
			false,             // no-wait
			nil,               // args
		)
		if err != nil {
			log.Printf("[rabbit-mq-error] [创建消费信道失败] %v\n", err)
		}

		// 保存监听和队列
		consumeQueueDeclares[queueDeclare.Name] = struct {
			queue    amqp.Queue
			messages <-chan amqp.Delivery
		}{
			queue:    queueDeclare,
			messages: receiveMessages,
		}
	}

	ctx, cancelFn = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	// 开始监听
	var forever chan struct{}
	if len(consumeQueueDeclares) > 0 {
		for _, consumeQueueDeclare := range consumeQueueDeclares {
			go RabbitMqConsume(consumeQueueDeclare)
		}
	}

	// receiveMessages, err := rabbitMqChannel.Consume(
	// 	queueDeclare.Name, // queue
	// 	"",                // consumer
	// 	true,              // auto-ack
	// 	false,             // exclusive
	// 	false,             // no-local
	// 	false,             // no-wait
	// 	nil,               // args
	// )
	// if err != nil {
	// 	log.Printf("[rabbit-mq-error] [创建消费信道失败] %v\n", err)
	// }

	// go func() {
	// 	for receiveMessage := range receiveMessages {
	// 		log.Printf("[rabbit-mq-error] [收到消息] %s\n", string(receiveMessage.Body))

	// 		business := &types.StdBusiness{}
	// 		if err = json.Unmarshal(receiveMessage.Body, business); err != nil {
	// 			log.Printf("[rabbit-mq-error] [解析业务失败] %v", err)
	// 		}

	// 		switch business.BusinessType {
	// 		case "ping":
	// 			log.Printf("[rabbit-mq-debug] [ping]")
	// 			// RabbitMqSendMessage(utils.NewCorrectWithBusiness("pong", "pong").Datum(map[string]any{"time": time.Now().Unix()}).ToJsonStr())
	// 		}
	// 	}
	// }()

	// log.Printf("[rabbit-mq-debug] [开始监听] %s\n", queueDeclare.Name)
	<-forever

}

func RabbitMqConsume(item struct {
	queue    amqp.Queue
	messages <-chan amqp.Delivery
}) {
	for receiveMessage := range item.messages {
		log.Printf("[rabbit-mq-error] [接收消息] %s\n", string(receiveMessage.Body))

		business := &types.StdBusiness{}
		if err = json.Unmarshal(receiveMessage.Body, business); err != nil {
			log.Printf("[rabbit-mq-error] [解析业务失败] %v", err)
		}

		switch business.BusinessType {
		case "ping":
			log.Printf("[rabbit-mq-debug] [ping]")
			// RabbitMqSendMessage(utils.NewCorrectWithBusiness("pong", "pong").Datum(map[string]any{"time": time.Now().Unix()}).ToJsonStr())
		}
	}
}
