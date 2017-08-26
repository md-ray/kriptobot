package service

import (
	"encoding/json"
	"log"
	"time"

	rmq "github.com/adjust/rmq"
)

var qInbound rmq.Queue
var qOutbound rmq.Queue

type InboundConsumer struct{}

func (consumer *InboundConsumer) Consume(delivery rmq.Delivery) {
	var inb BotMessage
	if err := json.Unmarshal([]byte(delivery.Payload()), &inb); err != nil {
		// handle error
		delivery.Reject()
		log.Println("error in inbound consumer")
		return
	}

	// perform task
	log.Printf("incoming message: %s ||dari %s ", inb, inb.Username)

	// prepare test reply
	var outb BotMessage
	outb.ChatID = inb.ChatID
	outb.Text = "ini balasan = " + inb.Text
	outbbytes, _ := json.Marshal(outb)
	qOutbound.PublishBytes(outbbytes)

	delivery.Ack()
}

func init() {
	connection := rmq.OpenConnection("queue-service", "tcp", "localhost:6379", 1)

	// inbound
	qInbound = connection.OpenQueue("telegram-inbound")
	inboundConsumer := &InboundConsumer{}
	qInbound.StartConsuming(10, time.Second)
	qInbound.AddConsumer("inbound consumer", inboundConsumer)

	// outbound
	qOutbound = connection.OpenQueue("telegram-outbound")

}
