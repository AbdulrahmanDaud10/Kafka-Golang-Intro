package main

import (
	kafkaConfig "github.com/AbdulrahmanDaud10/golang-kafka-intro/config"
	"github.com/AbdulrahmanDaud10/golang-kafka-intro/consumer"
)

func main() {
	topic := kafkaConfig.CONST_TOPIC
	consumer.Consume(topic)
}
