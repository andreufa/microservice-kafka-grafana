package main

import (
	"context"
	"log"
	"math"
	"processor-service/kafkalib"
	"processor-service/metrics"
	"time"
)

func main() {
	// Prometheus
	go func() {
		metrics.Listen("0.0.0.0:8083")
	}()
	ctx := context.Background()
	addreses := []string{"kafka:9092"}
	topic := "describe"
	groupId := "First group"
	clientId := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	consumer := kafkalib.NewConsumer(addreses, topic, groupId, clientId)
	defer consumer.Close()
	for {
		msgBytes, err := consumer.Read(ctx)
		if err != nil {
			log.Printf("Kafka error: %v. Reconnecting...", err)
			metrics.ObserveRequest(topic, err.Error())

			// 1. Переподключение с экспоненциальной задержкой
			time.Sleep(exponentialBackoff(5, 2*time.Second))
			continue
		}

		metrics.ObserveRequest(topic, "")
		log.Printf("Получено сообщение: %s\n", string(msgBytes))
	}
}

func exponentialBackoff(attempt int, maxDelay time.Duration) time.Duration {
	delay := time.Duration(math.Pow(2, float64(attempt))) * time.Second
	if delay > maxDelay {
		return maxDelay
	}
	return delay
}
