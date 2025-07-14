package main

import (
	"context"
	"log"
	"processor-service/kafkalib"
	"time"
)

func main() {
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
			continue
		}
		myMsg := string(msgBytes)
		log.Printf("Получено сообщение: %s\n", myMsg)
	}
}
