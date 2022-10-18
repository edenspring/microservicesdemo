package messages

import (
	"chatApp/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetAllMessages(c *gin.Context){
	var messages []db.Chat

	if result := h.DB.Find(&messages); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &messages)
}
