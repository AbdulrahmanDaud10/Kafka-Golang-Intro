package main

import (
	kafkaConfig "github.com/AbdulrahmanDaud10/golang-kafka-intro/config"
	"github.com/AbdulrahmanDaud10/golang-kafka-intro/producer"
)

func main() {
	topic := kafkaConfig.CONST_TOPIC
	producer.Produce(topic, 1000)
}
