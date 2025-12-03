package server

import (
	"Eino/internal/agent"
	"context"
	"fmt"
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

		// === 设置 SSE 流式返回 ===
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Flush()

		// === 流式运行 ===
		finalReply, err := ag.RunAgent(ctx, req.Query, func(chunk string) {
			// 每一个 token 回调时发送给前端
			fmt.Fprintf(c.Writer, "data: %s\n\n", chunk)
			c.Writer.Flush()
		})

		if err != nil {
			fmt.Fprintf(c.Writer, "data: [ERROR] %s\n\n", err.Error())
			c.Writer.Flush()
			return
		}

		// 最终结束（可选发送一个 [DONE]）
		fmt.Fprintf(c.Writer, "data: [DONE]\n\n")
		c.Writer.Flush()

		// （如果你想也可以写入 session 的最终回复，此时 finalReply 已经是完整回复）
		_ = finalReply
	})

	return router
}
