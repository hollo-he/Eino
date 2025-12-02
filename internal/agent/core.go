package agent

import (
	"Eino/internal/tools"
	"context"
	"io"
	"log"

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

// 工具agent
func (ag *Agent) ToolAgent(ctx context.Context, msg *schema.Message) (string, string, error) {
	systemPrompt := ag.SystemPrompt
	toolcall, err := ag.Model.Model.Generate(ctx, []*schema.Message{
		{
			Role:    schema.System,
			Content: systemPrompt,
		},
		msg,
	})
	if err != nil {
		return "", "", err
	}
	if len(toolcall.ToolCalls) == 0 {
		return "无需使用工具,直接回复", "", err
	}
	toolCall := toolcall.ToolCalls[0]
	args := toolCall.Function.Arguments
	name := toolCall.Function.Name

	//校验工具合法性
	if !ag.isToolAllowed(name) {
		rejectmsg := "Agent 不允许使用:" + name
		return rejectmsg, "", nil
	}

	out, err := tools.RunTool(ctx, name, args)
	if err != nil {
		out = "工具执行失败" + err.Error()
	}
	log.Println("工具输出:", out)
	return out, name, nil
}

// 带工具的
func (ag *Agent) RunAgent(ctx context.Context, msg string, onChunk func(string2 string)) (string, error) {

	user := schema.Message{Role: schema.User, Content: msg}
	GlobalSession.AddMessage(&user)

	systemPrompt := ag.SystemPrompt

	//消息体
	message := []*schema.Message{
		{Role: schema.System, Content: systemPrompt},
	}
	message = append(message, GlobalSession.Messages...)

	toolag := GlobalAgents["tool"]
	toolres, toolname, err := toolag.ToolAgent(ctx, &user)
	if err != nil {
		return "", err
	}

	//将结果与开始对话内容一起给大模型
	toolmsg := &schema.Message{Role: schema.Tool, Name: toolname, Content: toolres}

	message = append(message, toolmsg)
	GlobalSession.AddMessage(toolmsg)
	res, err := ag.Model.Model.Stream(ctx, message)
	if err != nil {
		return "", err
	}
	defer res.Close()
	var finalReply string

	for {
		chunk, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		text := chunk.Content
		finalReply += text

		// ⬅⬅⬅ 关键：将流式输出通过回调返回
		if onChunk != nil {
			onChunk(text)
		}
	}

	//////////////////////////////
	// 4) 写入 session
	//////////////////////////////
	GlobalSession.AddMessage(&schema.Message{
		Role:    schema.Assistant,
		Content: finalReply,
	})

	//////////////////////////////
	// 5) 最终 return（一般不会给前端，只是内部保存）
	//////////////////////////////
	return finalReply, nil
}

func (a *Agent) isToolAllowed(name string) bool {
	for _, t := range a.ToolNames {
		if t == name {
			return true
		}
	}
	return false
}
