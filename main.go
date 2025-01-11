package main

import (
	"laskerbot/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	botServer := gin.Default()

	routes.SendMove(botServer)
	routes.SendMessage(botServer)

	botServer.Run(":8080")

}
