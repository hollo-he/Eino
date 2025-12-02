package agent

import (
	"Eino/internal/llm"
	"Eino/internal/tools"
	"errors"
	"fmt"
	"os"

	"github.com/cloudwego/eino/schema"
)

var (
	ErrModelNotReady  = errors.New("llm model not initialized")
	ErrToolNotAllowed = errors.New("tool not allowed for agent")
	ErrToolNotFound   = errors.New("tool not found in registry")
)

type Agent struct {
	Name         string
	SystemPrompt string
	ToolNames    []string
	ToolInfos    []*schema.ToolInfo
	Model        *llm.Glm
}

var GlobalAgents = map[string]*Agent{}

// NewToolAgent toolagent构建(主要是工具的调用的选择)
func NewToolAgent(name string, toolNames []string) (*Agent, error) {

	//模型(脑子🧠)在不在
	if llm.GlmModel.Model == nil {
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

	//提示词获取
	promptPath := fmt.Sprintf("internal/agent/prompt/%sprompt.md", name)
	content, err := os.ReadFile(promptPath)
	if err != nil {
		return nil, fmt.Errorf("无法读取 %s 的 system prompt: %v", name, err)
	}

	ag := &Agent{
		Name:         name,
		SystemPrompt: string(content),
		ToolNames:    toolNames,
		ToolInfos:    toolinfos,
		Model:        llm.GlmModel,
	}

	if err := ag.Model.Model.BindTools(toolinfos); err != nil {
		return nil, err
	}
	return ag, nil
}

// NewChatAgent chatagent构建
func NewChatAgent(name string) (*Agent, error) {

	//模型(脑子🧠)在不在
	if llm.GlmModel.Model == nil {
		return nil, ErrModelNotReady
	}

	//提示词获取
	promptPath := fmt.Sprintf("internal/agent/prompt/%sprompt.md", name)
	content, err := os.ReadFile(promptPath)
	if err != nil {
		return nil, fmt.Errorf("无法读取 %s 的 system prompt: %v", name, err)
	}

	ag := &Agent{
		Name:         name,
		SystemPrompt: string(content),
		Model:        llm.GlmModel,
	}
	return ag, nil
}
