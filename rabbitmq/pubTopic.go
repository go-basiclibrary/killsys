package main

import (
	"fmt"
	"rabbitmq/mq"
)

func main() {
	mqOne := mq.NewRabbitMqTopic("newTopic", "wang.topic.one")
	mqTwo := mq.NewRabbitMqTopic("newTopic", "wang.topic.two")
	for i := 0; i <= 10; i++ {
		mqOne.PublishTopic(fmt.Sprintf("wangshao with topic one %d", i))
		mqTwo.PublishTopic(fmt.Sprintf("wangshao with topic two %d", i))
	}
}
