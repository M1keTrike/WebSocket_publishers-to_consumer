package main

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:8081/ws/handshake"
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error conectando:", err)
	}
	defer c.Close()

	var messages []string
	var mutex sync.Mutex

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Error leyendo:", err)
			return
		}

		mutex.Lock()
		messages = append(messages, string(message))
		currentLength := len(messages)
		mutex.Unlock()

		log.Printf("Mensaje recibido: %s", message)
		log.Printf("Total de mensajes almacenados: %d", currentLength)
	}
}