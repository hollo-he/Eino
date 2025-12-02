package llm

import (
	"Eino/internal/config"
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/model/openai"
)

type Glm struct {
	Model *openai.ChatModel
}

var GlmModel *Glm

func NewGlmModel() {

	ctx := context.Background()
	config.Load()
	chatmodel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  config.Cfg.ZhipuApiKey,
		BaseURL: config.Cfg.ZhipuAiUrl,
		Model:   config.Cfg.ZhipuAiName,
	})
	if err != nil {
		log.Println("NewChatModel failed, err=", err)
	}
	GlmModel = &Glm{
		Model: chatmodel,
	}
}
