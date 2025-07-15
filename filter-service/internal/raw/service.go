package raw

import (
	"context"
	"log"
	"time"
)

type RawService struct {
	rawRepository   *RawRepository
	kafkaRepository *KafkaRepository
}

func NewRawService(rawRepo *RawRepository, kafkaRepo *KafkaRepository) *RawService {
	return &RawService{
		rawRepository:   rawRepo,
		kafkaRepository: kafkaRepo,
	}
}

func (s *RawService) SaveToBD(raw, valid string) error {
	rawData := RawData{
		Data:  raw,
		Valid: valid,
	}

	if rawData.Valid == "true" {
		go func() {
			// 1. Отправка с ретраями и таймаутом
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			err := retryOperation(func() error {
				return s.kafkaRepository.Send(ctx, []byte(rawData.Data))
			}, 3, 1*time.Second) // 3 попытки с интервалом 1 сек

			if err != nil {
				// 2. Сохранение флага что сообщение не отправлено в кафку
				log.Printf("Failed to send to Kafka after retries: %v. Saving to DLQ...", err)
				// Далее действия с этими данными
			}
		}()

	}

	_, err := s.rawRepository.Save(&rawData)
	return err
}

func retryOperation(op func() error, maxRetries int, delay time.Duration) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		if err = op(); err == nil {
			return nil
		}
		time.Sleep(delay)
	}
	return err
}
