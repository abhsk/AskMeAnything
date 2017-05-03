package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/abhsk/AskMeAnything/telegram"

	"github.com/gorilla/mux"
)

const (
	BASEURL = "https://api.telegram.org/bot"
	APIKEY  = "351267960:AAFfApJepRNqFGGKaUcd30xkWoGFD2I118E"
)

func IndexHandler() http.HandlerFunc {
	//"message":{
	// "message_id":271,"from":{"id":35364348,"first_name":"Abhishek","last_name":"Pradhan"},
	// "chat":{"id":35364348,"first_name":"Abhishek","last_name":"Pradhan","type":"private"},
	// "date":1493803385,"text":"/ping","entities":[{"type":"bot_command","offset":0,"length":5}]}}
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to AskMeAnything")
	}
}

func MessageHandler(bot *telegram.Bot) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("err: ", err)
		}
		message := &telegram.Message{}
		err = json.Unmarshal(body, r)
		if err != nil {
			fmt.Println("err: ", err)
		}

		bot.Respond(message)
	}
}

func main() {
	bot := &telegram.Bot{
		URL: BASEURL + APIKEY,
	}

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", IndexHandler()).Methods("GET")
	r.HandleFunc("/message", MessageHandler(bot)).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
