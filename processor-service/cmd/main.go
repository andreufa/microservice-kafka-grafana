package main

import (
	"context"
	"log"
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
		time.Sleep(1 * time.Second)
		msgBytes, err := consumer.Read(ctx)
		if err != nil {
			log.Println(err)
			metrics.ObserveRequest(topic, err.Error())
			continue
		}
		myMsg := string(msgBytes)
		metrics.ObserveRequest(topic, "")
		log.Printf("Получено сообщение: %s\n", myMsg)
	}
}
