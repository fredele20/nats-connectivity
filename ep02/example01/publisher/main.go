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
		connection.Publish("intros", []byte("Hello World!"))
		count++
		fmt.Printf("sent message %v", count)
		time.Sleep(1 * time.Second)
	}
}