package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	connection, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}

	defer connection.Close()

	count := 0
	for {
		data := fmt.Sprintf("Hello World, count is %v", count)
		reply, err := connection.Request("intros", []byte(data), 500 * time.Millisecond)
		time.Sleep(1 * time.Second)
		if err != nil {
			log.Printf("error sending message count = %v, err: %v", count, err)
			continue
		}
		count++
		fmt.Printf("sent message %v, got reply %v", count, string(reply.Data))
	}
}