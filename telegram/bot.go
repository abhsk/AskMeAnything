package telegram

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	duck "github.com/ajanicij/goduckgo/goduckgo"
)

type Bot struct {
	URL string
}

func (b *Bot) Respond(message *Message) {
	chatId := strconv.Itoa(message.Chat.Id)

	matched, _ := regexp.MatchString("^(/)*echo ", message.Text)
	if matched == true {
		b.Reply(chatId, strings.SplitN(message.Text, "echo  ", 2)[1])
		return
	}

	matched, _ = regexp.MatchString("^(/)*ping ", message.Text)
	if matched == true {
		b.Reply(chatId, "Pong")
		return
	}

	matched, _ = regexp.MatchString("^(/)*search ", message.Text)
	if matched == true {
		b.Reply(chatId, Search(strings.SplitN(message.Text, "search ", 2)[1]))
		return
	}
	b.Reply(chatId, "We do not support this yet")
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

func Search(msg string) string {
	message, _ := duck.Query(msg)
	if len(message.RelatedTopics) > 0 {
		result := message.RelatedTopics[0]
		return message.Heading + " : " + result.Text
	}
	return "Not found"
}
