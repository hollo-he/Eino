package agent

import "github.com/cloudwego/eino/schema"

type Session struct {
	Messages   []*schema.Message
	MaxHistory int
}

func NewAgentSession() *Session {
	return &Session{
		Messages:   []*schema.Message{},
		MaxHistory: 20,
	}
}

func (s *Session) AddMessage(msg *schema.Message) {
	s.Messages = append(s.Messages, msg)

	if len(s.Messages) > s.MaxHistory {
		//TODO 考虑加入一个agent进行总结处理
		s.Messages = s.Messages[len(s.Messages)-s.MaxHistory:]
	}
}
