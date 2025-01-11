package routes

import (
	"encoding/json"
	"fmt"
	"laskerbot/controller"
	"laskerbot/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SendMove(server *gin.Engine) {
	var user_message models.Message
	server.POST("/move", func(ctx *gin.Context) {
		if err := ctx.ShouldBindJSON(&user_message); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		llm_response := controller.ControllerLLM(user_message.UserMessage)
		fmt.Printf("Tipo de llm_response: %T\n", llm_response)
		fmt.Printf("Conteúdo de llm_response: %+v\n", llm_response)

		if len(llm_response) == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Resposta do LLM vazia",
			})
			return
		}

		responseString := fmt.Sprintf("%v", llm_response[0])

		responseString = strings.TrimPrefix(responseString, "```json\n")
		responseString = strings.TrimSuffix(responseString, "\n```")
		responseString = strings.ReplaceAll(responseString, "```", "")

		var response struct {
			LLMResponse struct {
				AlgebricNote string `json:"algebric_note"`
				Explanation  string `json:"explanation"`
			} `json:"llm_response"`
		}

		if err := json.Unmarshal([]byte(responseString), &response); err != nil {
			fmt.Printf("Erro ao deserializar JSON: %s\n", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Erro ao processar resposta do LLM",
			})
			return
		}

		ctx.JSON(http.StatusOK, response)
	})
}
