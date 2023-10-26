package cmd

import (
	kafkaConfig "github.com/AbdulrahmanDaud10/golang-kafka-intro/config"
	"github.com/AbdulrahmanDaud10/golang-kafka-intro/consumer"
)

func Main() {
	topic := kafkaConfig.CONST_TOPIC
	consumer.Consume(topic)
}
