package service

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	rmq "github.com/adjust/rmq"
	"gopkg.in/telegram-bot-api.v4"
)

type OutboundConsumer struct{}

func (consumer *OutboundConsumer) Consume(delivery rmq.Delivery) {
	var outb BotMessage
	if err := json.Unmarshal([]byte(delivery.Payload()), &outb); err != nil {
		// handle error
		log.Println("error in outbound consumer")
		delivery.Reject()
		return
	}

	// perform task
	msg := tgbotapi.NewMessage(outb.ChatID, outb.Text)
	// msg.ReplyToMessageID = outb.MessageID
	fmt.Printf("preparing sending %s content %s\n", outb.ChatID, outb.Text)
	bot.Send(msg)
	delivery.Ack()
}

var bot *tgbotapi.BotAPI

func init() {
	var err error
	bot, err = tgbotapi.NewBotAPI("348636720:AAGRmH4FYDBPLxIH5qIoXRdUj5rvZjAII8A")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// handle outbound
	outboundConsumer := &OutboundConsumer{}
	qOutbound.StartConsuming(10, time.Second)
	qOutbound.AddConsumer("outbound consumer", outboundConsumer)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		var inboundMsg BotMessage
		inboundMsg.MessageID = update.Message.MessageID
		inboundMsg.ChatID = update.Message.Chat.ID
		inboundMsg.Date = update.Message.Date
		inboundMsg.Text = update.Message.Text
		inboundMsg.Uid = update.Message.From.ID
		inboundMsg.Username = update.Message.From.UserName
		inb, _ := json.Marshal(inboundMsg)
		qInbound.PublishBytes(inb)

		//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//		msg.ReplyToMessageID = update.Message.MessageID

		//		bot.Send(msg)
	}
}
