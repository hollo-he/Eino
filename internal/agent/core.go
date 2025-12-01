package agent

import (
	"Eino/internal/tools"
	"context"

	"github.com/cloudwego/eino/schema"
)

// 正常简单聊天,无工具
func (ag *Agent) Chat(ctx context.Context, msg string) (string, error) {
	if ag.Model == nil || ag.Model.Model == nil {
		return "", ErrModelNotReady
	}
	res, err := ag.Model.Model.Generate(ctx, []*schema.Message{
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
func (ag *Agent) RunAgent(ctx context.Context, msg string) (string, error) {

	user := schema.Message{Role: schema.User, Content: msg}
	GlobalSession.AddMessage(&user)

	systemPrompt := ag.SystemPrompt

	//消息体
	message := []*schema.Message{
		{Role: schema.System, Content: systemPrompt},
	}
	message = append(message, GlobalSession.Messages...)
	for {
		//提问->打算用的工具->开用->工具结果->大模型给出回复

		//调用工具需求
		res, err := ag.Model.Model.Generate(ctx, message)
		if err != nil {
			return "", err
		}
		if len(res.ToolCalls) == 0 {
			GlobalSession.AddMessage(res)
			return res.Content, nil
		}

		//工具输出
		toolCall := res.ToolCalls[0]
		args := toolCall.Function.Arguments
		name := toolCall.Function.Name

		//校验工具合法性
		if !ag.isToolAllowed(name) {
			rejectmsg := "Agent 不允许使用:" + name
			message = append(message,
				&schema.Message{Role: schema.Assistant, Content: "我试图调用工具：" + name},
				&schema.Message{Role: schema.Tool, Name: name, Content: rejectmsg},
			)
			continue
		}

		out, err := tools.RunTool(ctx, name, args)
		if err != nil {
			out = "工具执行失败" + err.Error()
		}

		//将结果与开始对话内容一起给大模型
		assistantToolCallMSG := &schema.Message{
			Role:      schema.Assistant,
			ToolCalls: res.ToolCalls,
		}
		toolmsg := &schema.Message{Role: schema.Tool, Name: toolCall.Function.Name, Content: out}

		message = append(message, assistantToolCallMSG, toolmsg)
		GlobalSession.AddMessage(assistantToolCallMSG)
		GlobalSession.AddMessage(toolmsg)
	}

}

func (a *Agent) isToolAllowed(name string) bool {
	for _, t := range a.ToolNames {
		if t == name {
			return true
		}
	}
	return false
}
