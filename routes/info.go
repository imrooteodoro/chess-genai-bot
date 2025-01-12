package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfoBot(server *gin.Engine) {

	about := `LaskerBot with Genai v1.0`
	server.GET("/about", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"about": about,
		})
	})

}
