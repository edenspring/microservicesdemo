package messages

import (
	"chatApp/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateMessageRequestBody struct {
	UserId    int    `json:"userId"`
	Content   string `json:"content"`
	ChannelId int    `json:"channelId"`
}

func (h handler) UpdateMessage(c *gin.Context) {
	body := UpdateMessageRequestBody{}
	id := c.Param("id")

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var message db.Chat

	if result := h.DB.First(&message, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	message.UserId = body.UserId
	message.Content = body.Content
	message.ChannelId = body.ChannelId

	h.DB.Save(&message)

	c.JSON(http.StatusCreated, &message)
}
