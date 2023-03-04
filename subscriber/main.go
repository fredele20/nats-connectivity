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
		log.Fatalf("can't connect to NATS %v", err)
	}

	defer connection.Close()

	connection.Subscribe("intros", func(msg *nats.Msg) {
		fmt.Printf("I got a message: %s\n", string(msg.Data))
	})

	time.Sleep(1 * time.Hour)
}