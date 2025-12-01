package agent

var GlobalSession *Session

func InitSession() {
	GlobalSession = NewAgentSession()
}
