package main

import (
	"Eino/internal/agent"
	"Eino/internal/llm"
	"Eino/internal/server"
	"Eino/internal/utils"
	"log"
)

// 主程序
func main() {
	//脑子🧠,启动!
	llm.NewOllamaModel()
	llm.NewGlmModel()

	//模式初始化
	searchTools := []string{"wikipedia_search", "duckduckgo_text_search"}
	toolAgent, err := agent.NewToolAgent("tool", searchTools)
	if err != nil {
		log.Fatalf("New searchAgent failed: %v", err)
	}
	chatAgent, err := agent.NewChatAgent("chat")

	agent.GlobalAgents["tool"] = toolAgent
	agent.GlobalAgents["default"] = chatAgent
	log.Println(agent.GlobalAgents)

	//会话初始化
	agent.InitSession()

	//神秘启动仪式
	r := server.New()
	log.Println("欢迎启动 Hollow 智能 Agent 🚀，监听端口 8080")
	utils.PrintBanner(`
██╗  ██╗ ██████╗ ██╗      ██╗      ██████╗ 
██║  ██║██╔═══██╗██║      ██║     ██╔═══██╗
███████║██║   ██║██║      ██║     ██║   ██║
██╔══██║██║   ██║██║      ██║     ██║   ██║
██║  ██║╚██████╔╝███████╗ ███████╗╚██████╔╝
╚═╝  ╚═╝ ╚═════╝ ╚══════╝ ╚══════╝ ╚═════╝ 

`)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
