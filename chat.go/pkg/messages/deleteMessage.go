package messages

import (
	"chatApp/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteMessage(c *gin.Context) {
	id := c.Param("id")

	var message db.Chat

	if result := h.DB.First(&message, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&message)

	c.Status(http.StatusOK)
}
