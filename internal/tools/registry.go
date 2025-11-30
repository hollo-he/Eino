package tools

import (
	"context"
	"errors"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

var toolMap = map[string]tool.InvokableTool{}
var toolInfos = []*schema.ToolInfo{}

func Register(toolInfo *schema.ToolInfo, toolInvokable tool.InvokableTool) {
	toolMap[toolInfo.Name] = toolInvokable
	toolInfos = append(toolInfos, toolInfo)
}

func RunTool(ctx context.Context, name, args string) (string, error) {
	if tool, ok := toolMap[name]; ok {
		return tool.InvokableRun(ctx, args)
	}
	return "", errors.New("调用不存在的工具:" + name)
}

func AllToolInfo() []*schema.ToolInfo {
	return toolInfos
}
func GetToolInfo(name string) (*schema.ToolInfo, error) {
	for _, toolInfo := range toolInfos {
		if toolInfo.Name == name {
			return toolInfo, nil
		}
	}
	return nil, errors.New("No Such Tool:" + name)
}
