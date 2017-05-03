package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome to AskMeAnything")
}

func messageHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//"message":{"message_id":271,"from":{"id":35364348,"first_name":"Abhishek","last_name":"Pradhan"},
	// "chat":{"id":35364348,"first_name":"Abhishek","last_name":"Pradhan","type":"private"},
	// "date":1493803385,"text":"/ping","entities":[{"type":"bot_command","offset":0,"length":5}]}}
}

func main() {
	router := httprouter.New()
	router.GET("/", indexHandler)
	router.POST("/message", messageHandler)

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
