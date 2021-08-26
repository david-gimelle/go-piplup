package main

import (
	"log"
	"net/http"
	"simple-websocket-chat/internal"
)

func main() {
	log.Println("Start listening message channel")
	go internal.ListenToMessageChannel()

	serverPort := "8080"
	log.Println("Starting web server on port ", serverPort)
	mux := internal.Routes()
	_ = http.ListenAndServe(":"+serverPort, mux)

}
