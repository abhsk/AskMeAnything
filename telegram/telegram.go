package telegram

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	SEND = "/sendmessage"
)

type Response struct {
	Message *Message
}

type Message struct {
	Id   int    `json:"message_id"`
	From User   `json:"from"`
	Chat Chat   `json:"chat"`
	Date int    `json:"date"`
	Text string `json:"text"`
}

type User struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type Chat struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

type Bot struct {
	URL string
}

func (b *Bot) Respond(message *Message) {
	chatId := strconv.Itoa(message.Chat.Id)
	msg := "Hello " + message.From.FirstName + "reply for " + message.Text

	url := b.URL + SEND + "?chat_id=" + chatId + "&text=" + msg
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err", err)
	}

	client := &http.Client{}
	_, e := client.Do(req)
	if e != nil {
		fmt.Println("e", e, url)
	}
}
