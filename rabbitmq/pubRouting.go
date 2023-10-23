package main

import (
	"fmt"
	"rabbitmq/mq"
)

func main() {
	mqOne := mq.NewRabbitMqRouting("newRouting", "route_one")
	mqTwo := mq.NewRabbitMqRouting("newRouting", "route_two")
	for i := 0; i <= 10; i++ {
		mqOne.PublishRouting(fmt.Sprintf("Hello wang shao one!msg is %d", i))
		mqTwo.PublishRouting(fmt.Sprintf("Hello wang shao two!msg is %d", i))
	}
}
