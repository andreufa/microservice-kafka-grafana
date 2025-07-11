package kafkapkg

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

const (
	partitionAny = -1 // Значение по умолчанию для PartitionAny
	flashTimeout = 500
	clientID     = "1001"
)

type Producer struct {
	producer *kafka.Writer
}

func NewProducer(address []string, topic string) *Producer {
	dialer := kafka.Dialer{
		ClientID: clientID,
	}
	conf := kafka.WriterConfig{
		Brokers:      address,
		Topic:        topic,
		Dialer:       &dialer,
		RequiredAcks: -1,    // Ждем подтверждения от всех реплик
		Async:        false, // Синхронная запись (ждем подтверждения)
		BatchTimeout: 0,     // Отправляем сразу, без батчинга
	}
	writer := kafka.NewWriter(conf)
	return &Producer{producer: writer}
}

func (p *Producer) Produce(message []byte) error {
	ctx := context.Background()
	err := p.producer.WriteMessages(ctx, kafka.Message{
		Value: message,
	})
	if err != nil {
		log.Printf("Ошибка при отправке сообщения в Produce(): %s", err)
		return err
	}
	return nil
}

func (p *Producer) Close() {
	p.producer.Close()
}

// 3. Важные параметры для надежной доставки
// RequiredAcks:

// kafka.RequireNone (0) - не ждем подтверждения

// kafka.RequireOne (1) - ждем подтверждения только от лидера

// kafka.RequireAll (-1) - ждем подтверждения от всех in-sync реплик (надежнее)

// Async:

// false - синхронная отправка (блокируемся до подтверждения)

// true - асинхронная отправка (не ждем подтверждения)

// Дополнительные параметры:

// go
// conf := kafka.WriterConfig{
//     // ...
//     MaxAttempts:     3,    // Количество попыток отправки
//     WriteTimeout:    10 * time.Second, // Таймаут записи
//     ReadTimeout:     10 * time.Second, // Таймаут чтения
//     BatchSize:       1,    // Отправлять по 1 сообщению (без батчинга)
