package main

import "rabbitmq/mq"

func main() {
	//获取所有的消息
	mqTopic := mq.NewRabbitMqTopic("newTopic", "#")
	mqTopic.ReceiveTopic()
}
