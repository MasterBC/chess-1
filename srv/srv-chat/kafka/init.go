package kafka

import (
pb "chess/srv/srv-chat/proto"
"github.com/Shopify/sarama"
cli "gopkg.in/urfave/cli.v2"
)

var (
    kAsyncProducer sarama.AsyncProducer
    kClient        sarama.Client
    ChatTopic      string
)

func initKafka(c *cli.Context) {
    addrs := c.StringSlice("kafka-brokers")
    ChatTopic = c.String("chat-topic")
    config := sarama.NewConfig()
    config.Producer.Return.Successes = false
    config.Producer.Return.Errors = false
    producer, err := sarama.NewAsyncProducer(addrs, config)
    if err != nil {
	panic(err)
    }

    kAsyncProducer = producer
    cli, err := sarama.NewClient(addrs, nil)
    if err != nil {
	panic(err)
    }
    kClient = cli
}

func Init(c *cli.Context) {
    initKafka(c)
}
func NewConsumer() (sarama.Consumer, error) {
    return sarama.NewConsumerFromClient(kClient)
}

func Input(message *pb.Chat_Message) {
    msg := &sarama.ProducerMessage{Topic: ChatTopic, Key: sarama.ByteEncoder(message.Id), Value: sarama.ByteEncoder(message.Body)}
    kAsyncProducer.Input() <- msg
}
