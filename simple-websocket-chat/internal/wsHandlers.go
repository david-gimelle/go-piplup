package internal

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var messagesChan = make(chan wsPayload)
var connections []*websocket.Conn

type wsPayload struct {
	Action     string `json:"action"`
	NewMessage string `json:"new_message"`
	UserName   string `json:"username"`
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketRender(w http.ResponseWriter, r *http.Request) {
	log.Println("Ws handler")
	wscon, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket Upgrade error", err)
	}
	log.Printf("Connected to a websocket with socket %p\n", wscon)

	connections = append(connections, wscon)
	log.Println("Number of connections registered", len(connections))

	go listenToWs(wscon)

}

func listenToWs(c *websocket.Conn) {
	log.Println("Start listening")

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Listening Routine for socket %p has been stoped", &c)
		}
	}()
	for {
		log.Println("In the For loop")
		var payload wsPayload
		err := c.ReadJSON(&payload)
		log.Printf("message received in ws %v", payload)

		if err != nil {
			log.Printf("Json WS Error %v\n", err)
		}

		messagesChan <- payload
	}

}

func ListenToMessageChannel() {
	for {
		m := <-messagesChan
		log.Printf("Message received in channel: %v\n", m)

		switch m.Action {

		case "PUBLISH":
			log.Printf("Will Broadcast to all this message %v", m.NewMessage)
			for _, conn := range connections {
				m.Action = "BROADCAST"
				conn.WriteJSON(m)
				fmt.Printf("Message sent to browser with conn %p \n", conn)
			}
		}
	}
}
