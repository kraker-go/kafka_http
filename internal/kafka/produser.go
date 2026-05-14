package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	write *kafka.Writer
}

func NewKafkaProducer() *KafkaProducer {
	return &KafkaProducer{
		write: &kafka.Writer{
			Addr:     kafka.TCP("127.0.0.1:9092"),
			Topic:    "movies",
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *KafkaProducer) Send(ctx context.Context, message []byte) error {
	return p.write.WriteMessages(ctx, kafka.Message{Value: message})
}
