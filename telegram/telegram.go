package telegram

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
