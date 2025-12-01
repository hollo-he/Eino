package server

import (
	"Eino/internal/agent"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AgentRequest struct {
	AgentName string `json:"agentName"`
	Query     string `json:"query"`
}

type AgentResponse struct {
	Answer string `json:"answer"`
}

func New() *gin.Engine {

	router := gin.Default()

	router.POST("/agent/:Agentname/run", func(c *gin.Context) {
		var req AgentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := context.Background()
		agName := c.Param("Agentname")
		log.Printf("Agent name: %s", agName)
		if agName == "" {
			agName = "default"
		}
		ag, ok := agent.GlobalAgents[agName]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "agent not found: " + agName})
			return
		}
		answer, err := ag.RunAgent(ctx, req.Query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//utils.Edge_tts(answer)

		//go utils.Win_tts(answer)

		c.JSON(http.StatusOK, gin.H{"answer": answer})

	})

	return router
}
