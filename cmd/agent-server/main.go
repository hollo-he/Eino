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

	//模式初始化
	searchTools := []string{"wikipedia_search", "duckduckgo_text_search"}
	searchAgent, err := agent.NewAgent("search", searchTools)
	if err != nil {
		log.Fatalf("New searchAgent failed: %v", err)
	}
	agent.GlobalAgents["wikipedia_search"] = searchAgent
	agent.GlobalAgents["default"] = searchAgent

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
