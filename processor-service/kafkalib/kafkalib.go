package kafkalib

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(address []string, topic, groupId, clientId string) *Consumer {
	dialer := &kafka.Dialer{
		ClientID: clientId, // Уникальный ID для этого потребителя
	}
	conf := kafka.ReaderConfig{
		Brokers: address,
		Topic:   topic,
		GroupID: groupId,
		Dialer:  dialer,
	}
	read := kafka.NewReader(conf)
	return &Consumer{
		reader: read,
	}
}

func (c *Consumer) Read(ctx context.Context) ([]byte, error) {
	msg, err := c.reader.ReadMessage(ctx)
	if err != nil {
		log.Printf("Ошибка при получении сообщения Read() :", err)
		return nil, err
	}
	return msg.Value, nil

}

func (c *Consumer) Close() {
	c.reader.Close()
}
