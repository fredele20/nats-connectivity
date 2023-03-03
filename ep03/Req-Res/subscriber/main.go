package main

import (
	"encoding/json"
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
		pl := &model.Payload{}
		json.Unmarshal(msg.Data, pl)
		replyData := fmt.Sprintf("ack message # %v", pl.Count)
		msg.Respond([]byte(replyData))
		fmt.Printf("I got a message: %s\n", pl.Data)
	})

	time.Sleep(1 * time.Hour)
}
