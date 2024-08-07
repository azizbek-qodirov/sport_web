package handlers

import (
	"fmt"
	"net/http"
	"project/models"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) GetMatches(c *gin.Context) {
	matches, err := h.Match.GetMatches()
	// pp.Println(matches)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get matches", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, matches)
}

func (h *HTTPHandler) GetMatchDetails(c *gin.Context) {
	matchID := c.Param("id")

	// Get match details
	matches, err := h.Match.GetMatchByID(matchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get match details", "details": err.Error()})
		return
	}

	// Get chat messages
	messages, err := h.Chat.GetMessages(matchID)
	fmt.Println(matchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages", "details": err.Error()})
		return
	}

	response := models.MatchDetailsResponse{
		Match:    matches,
		Messages: messages,
	}

	c.JSON(http.StatusOK, response)
}
