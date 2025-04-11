package main

import (
	"web_socket/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	infrastructure.Routes(r)

	r.Run(":8081")

}
