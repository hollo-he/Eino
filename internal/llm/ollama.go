package llm

import (
	"Eino/internal/config"
	"Eino/internal/tools"
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/schema"
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

// 正常简单聊天,无工具
func (o *Ollama) Chat(ctx context.Context, msg string) (string, error) {
	res, err := o.Model.Generate(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: msg,
		},
	})
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

// 带工具的
func (o *Ollama) RunAgent(ctx context.Context, msg string) (string, error) {

	//消息体
	message := []*schema.Message{
		{
			//系统提示词
			Role: schema.System,
			Content: "你是一个 ReAct Agent。用户每次提问你必须使用工具!" +
				"你可以使用以下工具：维基百科(wikipedia_search)、DuckDuckGo(duckduckgo_text_search)。",
		},
		{
			//用户提示词
			Role:    schema.User,
			Content: msg,
		},
	}
	for {
		//提问->打算用的工具->开用->工具结果->大模型给出回复

		//调用工具需求
		res, err := o.Model.Generate(ctx, message)
		if err != nil {
			return "", err
		}
		if len(res.ToolCalls) == 0 {
			return res.String(), nil
		}

		//工具输出
		toolCall := res.ToolCalls[0]
		args := toolCall.Function.Arguments
		name := toolCall.Function.Name

		out, err := tools.RunTool(ctx, name, args)
		if err != nil {
			return "", err
		}

		//将结果与开始对话内容一起给大模型

		message = append(message, &schema.Message{
			Role:      schema.Assistant,
			ToolCalls: res.ToolCalls,
		},
			&schema.Message{
				Role:    schema.Tool,
				Name:    toolCall.Function.Name,
				Content: out,
			})
	}
}
