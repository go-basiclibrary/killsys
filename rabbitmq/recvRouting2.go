package main

import "rabbitmq/mq"

func main() {
	mqOne := mq.NewRabbitMqRouting("newRouting", "route_two")
	mqOne.ReceiveRouting()
}
