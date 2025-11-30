package agent

import (
	"Eino/internal/llm"
	"Eino/internal/tools"
	"errors"

	"github.com/cloudwego/eino/schema"
)

var (
	ErrModelNotReady  = errors.New("llm model not initialized")
	ErrToolNotAllowed = errors.New("tool not allowed for agent")
	ErrToolNotFound   = errors.New("tool not found in registry")
)

type Agent struct {
	Name      string
	ToolNames []string
	ToolInfos []*schema.ToolInfo
	Model     *llm.Ollama
}

var GlobalAgents = map[string]*Agent{}

// agent构建(主要是工具的调用的选择)
func NewAgent(name string, toolNames []string) (*Agent, error) {

	//模型(脑子🧠)在不在
	if llm.OllamaChatModel.Model == nil {
		return nil, ErrModelNotReady
	}

	//工具校验与注册
	toolinfos := []*schema.ToolInfo{}
	for _, toolnames := range toolNames {
		ti, err := tools.GetToolInfo(toolnames)
		if err != nil {
			return nil, err
		}
		toolinfos = append(toolinfos, ti)
	}

	ag := &Agent{
		Name:      name,
		ToolNames: toolNames,
		ToolInfos: toolinfos,
		Model:     llm.OllamaChatModel,
	}

	if err := ag.Model.Model.BindTools(toolinfos); err != nil {
		return nil, err
	}

	return ag, nil
}
