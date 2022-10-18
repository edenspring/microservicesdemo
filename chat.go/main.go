package main

import (
	"chatApp/pkg/db"
	"chatApp/pkg/messages"
	"github.com/gin-gonic/gin"
	) 

func main() {
	go h.run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	handler := db.Init()

	messages.RegisterRoutes(router, handler)
	router.Run("0.0.0.0:8080")
}

