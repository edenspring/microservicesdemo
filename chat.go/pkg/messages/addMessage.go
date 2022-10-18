package messages

import (
	"chatApp/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddMessageRequestBody struct {
	UserId    int    `json:"userId"`
	Content   string `json:"content"`
	ChannelId int    `json:"channelId"`
}

func (h handler) AddMessage(c *gin.Context) {
	body := AddMessageRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var message db.Chat

	message.UserId = body.UserId
	message.Content = body.Content
	message.ChannelId = body.ChannelId

	if result := h.DB.Create(&message); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &message)
}
