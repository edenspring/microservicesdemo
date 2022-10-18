package messages

import	(
	"chatApp/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetChannelMessages(c *gin.Context){
	channelId := c.Param("id")

	var messages []db.Chat

	if result := h.DB.Where("channel_id = ?", channelId).Find(&messages); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &messages)
}
