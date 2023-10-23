package main

import "rabbitmq/mq"

func main() {
	mq := mq.NewRabbitMqSimple("simpleQueue")
	mq.PublishSimple("hello,wangshao!GoodAfternoon")
}
