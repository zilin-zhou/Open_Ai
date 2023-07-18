package main

import (
	"open_ai_chat/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouters(r)
	r.Run("127.0.0.1:8080")
}
