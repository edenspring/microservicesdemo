package messages

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/messages")
	routes.POST("/", h.AddMessage)
	routes.GET("/", h.GetAllMessages)
	routes.GET("/:id", h.GetChannelMessages)
	routes.PUT("/:id", h.UpdateMessage)
	routes.DELETE("/:id", h.DeleteMessage)
}