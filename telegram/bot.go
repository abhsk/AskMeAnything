package telegram

import (
	"fmt"
	"net/http"
	"strconv"
)

type Bot struct {
	URL string
}

func (b *Bot) Respond(message *Message) {
	chatId := strconv.Itoa(message.Chat.Id)
	msg := "Hello " + message.From.FirstName + "reply for " + message.Text

	b.Reply(chatId, msg)
}

func (b *Bot) Reply(chatId, message string) {
	url := b.URL + SEND + "?chat_id=" + chatId + "&text=" + message
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
