package server

import (
	"Eino/internal/llm"
	"Eino/internal/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AgentRequest struct {
	Query string `json:"query"`
}

type AgentResponse struct {
	Answer string `json:"answer"`
}

func New() *gin.Engine {

	router := gin.Default()

	router.POST("/agent/run", func(c *gin.Context) {
		var req AgentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := context.Background()
		answer, err := llm.OllamaChatModel.RunAgent(ctx, req.Query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//utils.Edge_tts(res)

		go utils.Win_tts(answer)

		c.JSON(http.StatusOK, gin.H{"answer": answer})

	})

	return router
}
