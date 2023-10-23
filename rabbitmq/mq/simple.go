package mq

import (
	"fmt"
	tlog "git.tencent.com/trpc-go/trpc-go/log"
	"log"

	"github.com/streadway/amqp"
)

const mqURL = "amqp://wangShao:wangShao@106.52.228.207:5673/wangShao"

type RabbitMq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// Key
	Key string
	// 连接信息
	MqURL string
}

// NewRabbitMq 创建Mq实例
func NewRabbitMq(queueName string, exchange string, key string, mqURL string) *RabbitMq {
	rabbitMq := &RabbitMq{QueueName: queueName, Exchange: exchange, Key: key, MqURL: mqURL}

	var err error
	rabbitMq.conn, err = amqp.Dial(rabbitMq.MqURL)
	rabbitMq.failOnErr(err, "create conn err!")
	rabbitMq.channel, err = rabbitMq.conn.Channel()
	rabbitMq.failOnErr(err, "get channel failed!")

	return rabbitMq
}

// Destory 销毁Mq实例
func (mq *RabbitMq) Destory() {
	mq.conn.Close()
	mq.channel.Close()
}

// failOnErr err detail
func (mq *RabbitMq) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// NewRabbitMqSimple queueName队列名称
func NewRabbitMqSimple(queueName string) *RabbitMq {
	// exchange为""，默认使用default
	return NewRabbitMq(queueName, "", "", mqURL)
}

// PublishSimple Simple模式推送消息  生产者
func (mq *RabbitMq) PublishSimple(message string) {
	//申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	//保证队列存在，消息能发送到队列中
	_, err := mq.channel.QueueDeclare(
		mq.QueueName, // 队列名称
		false,        // 是否持久化
		false,        // 是否为自动删除
		false,        // 是否具有排他性，不常用，创建一个自己可见的队列，其他用户不可见
		false,        // 是否阻塞，发送消息后是否等待响应
		nil,          // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}
	// 发送消息到队列中
	err = mq.channel.Publish(
		mq.Exchange,  // 为""则使用default队列
		mq.QueueName, // 队列名称
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

// ConsumeSimple Simple模式接收消息  消费者
func (mq *RabbitMq) ConsumeSimple() {
	//申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	//保证队列存在，消息能发送到队列中
	_, err := mq.channel.QueueDeclare(
		mq.QueueName, // 队列名称
		false,        // 是否持久化
		false,        // 是否为自动删除
		false,        // 是否具有排他性，不常用，创建一个自己可见的队列，其他用户不可见
		false,        // 是否阻塞，发送消息后是否等待响应
		nil,          // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}
	// 发送消息到队列中
	msgs, err := mq.channel.Consume(
		mq.QueueName, // 为""则使用default队列
		"",           // 用来区分多个消费者
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
