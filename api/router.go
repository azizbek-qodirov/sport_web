package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project/api/handlers"
	"project/storage"
)

func NewRouter(dbs *storage.Storage) *gin.Engine {
	router := gin.Default()
	h := handlers.NewHandler(dbs)
	router.Use(CORSMiddleware())

	router.GET("/ws", h.HandleWSConnection)
	router.POST("/send", h.PostMessage)
	router.GET("/match/:id", h.GetMatchDetails)
	router.GET("/matches", h.GetMatches)

	go h.HandleMessages()

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
