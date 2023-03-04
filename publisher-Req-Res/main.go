package main

import (
	"awesome/model"
	"encoding/json"
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
	pl := &model.Payload{
		Data: "Hello World",
	}
	for {
		pl.Count = count
		data, _ := json.Marshal(pl)
		reply, err := connection.Request("intros", []byte(data), 500*time.Millisecond)
		time.Sleep(1 * time.Second)
		if err != nil {
			log.Printf("error sending message count = %v, err: %v\n", count, err)
			continue
		}
		count++
		pl := &model.Payload{data, count}
		json.Unmarshal(reply.Data, pl)
		fmt.Println()
		fmt.Printf("sent message %v, got reply %v\n", count, string(reply.Data))
	}
}
