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

var OllamaChatModel *Ollama

func NewOllamaModel() {

	//配置设置
	ctx := context.Background()
	config.Load()
	model, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: config.Cfg.OllamaUrl,
		Model:   config.Cfg.OllamaModelName,
	})
	if err != nil {
		log.Printf("NewChatModel failed, err=%v\n", err)
	}

	//整合,创建工具箱
	tools.AllToolInit()
	toolInfo := tools.AllToolInfo()
	if err := model.BindTools(toolInfo); err != nil {
		log.Printf("model bind tools failed, err=%v\n", err)
	}

	OllamaChatModel = &Ollama{
		Model: model,
	}

}
