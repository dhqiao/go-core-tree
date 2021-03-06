package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"fmt"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
//This is just an object with methods for taking a normal HTTP connection and upgrading
// it to a WebSocket as we'll see later in the code.
var upgrader = websocket.Upgrader{}

func main() {
	//......................................
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	log.Println("http server started on :1234")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

type Message struct {
	Message string `json:"message"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	clients[ws] = true
	for {
		var msg Message
		fmt.Println("---------")
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}
func handleMessages() {
	for {
		fmt.Println("===========")
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		msg.Message = msg.Message + "aaaa"
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
