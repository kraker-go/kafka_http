package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaConsumer struct {
	read *kafka.Reader
}

func NewKafkaConsumer() *KafkaConsumer {
	return &KafkaConsumer{
		read: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   "movies",
			GroupID: "movie-group",
		}),
	}
}
func (k *KafkaConsumer) Start(ctx context.Context, logg *zap.Logger) error {
	for {
		msg, err := k.read.ReadMessage(ctx)
		if err != nil {
			return fmt.Errorf("kafka: consumer - err", err)
		}
		logg.Info("Получено сообщение")
		zap.String("value", string(msg.Value))
	}

	return nil
}
