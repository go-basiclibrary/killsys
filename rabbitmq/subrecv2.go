package main

import "rabbitmq/mq"

func main() {
	mq := mq.NewRabbitMqSubScribe("newSubPub")
	mq.ReceiveSub()
}
