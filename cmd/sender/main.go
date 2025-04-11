package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:8081/ws/handshake"
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error conectando:", err)
	}
	defer c.Close()

	clientID := time.Now().UnixNano() % 1000
	counter := 0

	for {
		message := fmt.Sprintf("Cliente %d - Mensaje %d", clientID, counter)
		err := c.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Error escribiendo:", err)
			return
		}
		counter++
		time.Sleep(2 * time.Second)
	}
}