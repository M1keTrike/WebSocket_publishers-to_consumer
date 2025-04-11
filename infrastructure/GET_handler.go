package infrastructure

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WebSocketHandler(ctx *gin.Context) {

	var upg = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upg.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Printf("error de conexion: %v", err)
		return
	}

	defer conn.Close()

	log.Println("Cliente conectado")

	for {

		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Printf("error en la lectura %v", err)
			break
		}

		log.Printf("Recibido: %s, de tipo %d", p, messageType)

		conn.WriteMessage(messageType, p)

		time.Sleep(16 * time.Millisecond)

	}
}
