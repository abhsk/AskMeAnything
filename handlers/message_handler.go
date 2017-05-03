package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/abhsk/AskMeAnything/telegram"
)

func MessageHandler(bot *telegram.Bot) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("err: ", err)
		}
		fmt.Println("body: ", string(body))
		response := &telegram.Response{}
		err = json.Unmarshal(body, response)
		if err != nil {
			fmt.Println("err: ", err)
		}
		fmt.Println("message: ", response)

		bot.Respond(response.Message)
	}
}
