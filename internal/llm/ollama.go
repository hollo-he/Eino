package llm

import (
	"Eino/internal/config"
	"Eino/internal/tools"
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/model/ollama"
)

// ollama模型初始化
type Ollama struct {
	Model *ollama.ChatModel
}

var OllamaToolModel *Ollama
var OllamaChatModel *Ollama

func NewOllamaModel() {

	//配置设置
	ctx := context.Background()
	config.Load()
	toolmodel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: config.Cfg.OllamaUrl,
		Model:   config.Cfg.OllamaToolModelName,
	})
	if err != nil {
		log.Printf("NewToolModel failed, err=%v\n", err)
	}
	chatmodel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: config.Cfg.OllamaUrl,
		Model:   config.Cfg.OllamaChatModelName,
	})
	if err != nil {
		log.Println("NewChatModel failed, err=", err)
	}

	//整合,创建工具箱
	tools.AllToolInit()
	toolInfo := tools.AllToolInfo()
	if err := toolmodel.BindTools(toolInfo); err != nil {
		log.Printf("model bind tools failed, err=%v\n", err)
	}

	OllamaToolModel = &Ollama{
		Model: toolmodel,
	}
	OllamaChatModel = &Ollama{
		Model: chatmodel,
	}
}
