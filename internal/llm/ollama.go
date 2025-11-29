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

	//工具配置

	//维基百科📚
	tools.WikipediaInit()
	wikiinfo, err := tools.WikipediaTool.Info(ctx)
	if err != nil {
		log.Printf("Wikipedia tool info failed, err=%v\n", err)
	}

	//duck搜索🔍
	tools.DuckDuckGoInit()
	duckduckgo, err := tools.DuckDuckGo.Info(ctx)
	if err != nil {
		log.Printf("DuckDuckGo info failed, err=%v\n", err)
	}

	//整合,创建工具箱
	toolInfo := []*schema.ToolInfo{wikiinfo, duckduckgo}
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
			//用户提示词
			Role:    schema.User,
			Content: msg,
		},
		{
			//系统提示词
			Role:    schema.System,
			Content: "请务必先使用符合条件的工具后,再回答用户问题,不管问题多简单,你能使用的工具有wikipedia_search跟duckduckgo_search,一个是维基百科查询,一个是搜索引擎获取网页链接",
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
		toolCall := res.ToolCalls[0]

		//工具输出
		var toolOutput string
		switch toolCall.Function.Name {
		case "wikipedia_search":
			out, err := tools.WikipediaTool.InvokableRun(ctx, toolCall.Function.Arguments)
			if err != nil {
				return "", err
			}
			toolOutput = out
		case "duckduckgo_text_search":
			out, err := tools.DuckDuckGo.InvokableRun(ctx, toolCall.Function.Arguments)
			if err != nil {
				return "", err
			}
			toolOutput = out
		default:
			toolOutput = res.String()
		}

		//将结果与开始对话内容一起给大模型
		message = append(message, &schema.Message{
			Role:      schema.Assistant,
			ToolCalls: res.ToolCalls,
		},
			&schema.Message{
				Role:    schema.Tool,
				Name:    toolCall.Function.Name,
				Content: toolOutput,
			})
	}
}
