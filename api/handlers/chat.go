package handlers

import (
	"net/http"
	"project/models"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) PostMessage(c *gin.Context) {
	var message models.MessageCReq
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't bind request", "details": err.Error()})
		return
	}
	if err := h.Chat.CreateMessage(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create message", "details": err.Error()})
		return
	}
	broadcast <- message

	c.JSON(http.StatusCreated, gin.H{"message": "Message created successfully"})
}
