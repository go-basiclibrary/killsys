package main

import (
	"fmt"
	"rabbitmq/mq"
)

func main() {
	mq := mq.NewRabbitMqSubScribe("newSubPub")
	for i := 1; i <= 20; i++ {
		mq.PublishPub(fmt.Sprintf("订阅模式生成第%d条数据", i))
	}
}
