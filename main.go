package main

import (
	"log"
	"net/http"
	"os"

	h "github.com/abhsk/AskMeAnything/handlers"
	t "github.com/abhsk/AskMeAnything/telegram"

	"github.com/gorilla/mux"
)

const (
	BASEURL = "https://api.telegram.org/bot"
	APIKEY  = "351267960:AAFfApJepRNqFGGKaUcd30xkWoGFD2I118E"
)

func main() {
	bot := &t.Bot{
		URL: BASEURL + APIKEY,
	}

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", h.IndexHandler()).Methods("GET")
	r.HandleFunc("/message", h.MessageHandler(bot)).Methods("POST")

	// Bind to a port and pass our router in
	port := ":" + os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, r))
}
