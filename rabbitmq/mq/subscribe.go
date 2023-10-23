package mq

import (
	"fmt"
	tlog "git.tencent.com/trpc-go/trpc-go/log"
	"github.com/streadway/amqp"
)

// NewRabbitMqSubScribe exchange交换机名称
func NewRabbitMqSubScribe(exchange string) *RabbitMq {
	// queueName为""
	return NewRabbitMq("", exchange, "", mqURL)
}

// PublishPub 订阅模式生产
func (mq *RabbitMq) PublishPub(message string) {
	//创建交换机
	err := mq.channel.ExchangeDeclare(
		mq.Exchange,
		// fanout广播类型
		"fanout",
		// 是否持久化
		true,
		false,
		// true表示这个exchange不可以被client用来推送消息,用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// 发送消息到队列中
	err = mq.channel.Publish(
		mq.Exchange, // 订阅模式需要指定Exchange
		"",          // 队列名称
		// 如果为true，根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		// 如果为true,当exchange发送消息到队列后发现没有绑定消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		tlog.Errorf("mq simple publish err: %s", err)
	}
}

// ReceiveSub 订阅模式消费
func (mq *RabbitMq) ReceiveSub() {
	//创建交换机
	err := mq.channel.ExchangeDeclare(
		mq.Exchange,
		// fanout广播类型
		"fanout",
		// 是否持久化
		true,
		false,
		// true表示这个exchange不可以被client用来推送消息,用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 创建队列
	q, err := mq.channel.QueueDeclare(
		"", // 随机生成队列
		false,
		false,
		true,
		false, nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 队列绑定exchange
	err = mq.channel.QueueBind(
		q.Name,
		"", // 订阅模式下key必须为空
		mq.Exchange,
		false, nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	// 发送消息到队列中
	msgs, err := mq.channel.Consume(
		q.Name, // 队列的名称不指定,随机生成
		"",     // 用来区分多个消费者
		// 是否自动应答,或者false时实现回调
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		false,
		nil,
	)
	if err != nil {
		tlog.Errorf("mq simple consume err: %s", err)
	}

	forever := make(chan struct{})
	// 协程处理消息
	go func() {
		tlog.Info("goroutine is start...")
		for d := range msgs {
			// 实现我们要处理的逻辑函数
			fmt.Printf("received msg:%s\n", string(d.Body))
		}
	}()

	// 阻塞
	<-forever
}
