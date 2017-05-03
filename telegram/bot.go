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

	request_message := strings.TrimSpace(message.Text)
	response_message := "We do not support this yet"

	matched, _ := regexp.MatchString("^(/)*echo", request_message)
	if matched == true {
		response_message = strings.SplitN(request_message, "echo", 2)[1]
	}

	matched, _ = regexp.MatchString("^(/)*ping", request_message)
	if matched == true {
		response_message = "pong"
	}

	matched, _ = regexp.MatchString("^(/)*search", request_message)
	if matched == true {
		response_message = Search(strings.SplitN(request_message, "search", 2)[1])
	}

	b.Reply(chatId, response_message)
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
