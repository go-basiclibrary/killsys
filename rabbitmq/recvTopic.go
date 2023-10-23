package main

import "rabbitmq/mq"

func main() {
	mq := mq.NewRabbitMqTopic("newTopic", "wang.*.two")
	mq.ReceiveTopic()
}
