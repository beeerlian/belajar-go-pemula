package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/beeerlian/websocket_gorilla/websocket"
)

func main() {
	fmt.Println("Go websocket")
	handleRouter()
}

func handleRouter() {
	http.HandleFunc("/ws", wsHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func wsHandle(w http.ResponseWriter, r *http.Request) {
	wsconn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sending msg to %v ", wsconn.LocalAddr())
	websocket.Writer(wsconn)
}
