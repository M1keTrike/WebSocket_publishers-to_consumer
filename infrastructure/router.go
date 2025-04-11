package infrastructure

import "github.com/gin-gonic/gin"

func Routes(en *gin.Engine) {
	wsg := en.Group("ws")
	{
		wsg.GET("handshake", WebSocketHandler)
	}
}
