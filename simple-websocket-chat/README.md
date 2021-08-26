# Simple chat example in Go with Gorilla Websocket

This go application shows how to build a simple chat application using websockets.

## How to use this application
- Start the application with this command
  ```go run ./cmd/main.go```
- Open this [http://127.0.0.1:8080](http://127.0.0.1:8080) in 2 different browsers or tabs
- Use the form to submit username and messages. The messages will be displayed in both tabs thanks to gorilla websockets.

## What libraries are used?
- github.com/CloudyKit/jet/v6 for the webpages in go
- net/http is a native network golang lib
- github.com/gorilla/websocket for websocket in go
- Bootstrap for Javascript

## How the source code is organised?
- cmd/main.go contains the main application
- html/home.jet contains the webpage with the chat form
- internal folder contains, routing, http handlers, websockets handlers and go routines listening to sockets

## How it works?
The application starts and listens
- The app Listen on port 8080 for http request
- A go routine listen on the go channel ```messagesChan``` with the line ```go internal.ListenToMessageChannel()```

When an user load a page:
- When the main page has been loaded, a DOM lister ```document.addEventListener("DOMContentLoaded" ...``` will make a call to ```https://127.0.0.1/ws``` to open a websocket with the go app
- The go application will use a ```websocket.Upgrader``` to upgrade the connection from HTTP to WebSocket Protocol. It also starts a go routine to listen to any incoming communication on this websocket: ```go listenToWs(wscon)```

From then, each time that "Enter" is taped in the message field:
- This is captured by the browser in the ```newMessageField.addEventListener("change", ...``` and a message with a json payload is send to the go server
- The go routine listening on this websocket receives this payload and send a message for broadcasting it to the go channel ```messagesChan```
- The message is received by the routine ```ListenToMessageChannel``` and the message is broadcasted to any existing websockets with ```conn.WriteJSON(m)```
- Each browser/tab receive the message by websocket and catch with the event ```socket.onmessage = message ..``` that displays the message in the chat section of the page
