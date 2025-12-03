package tools

import "fmt"

func AllToolInit() {

	var toolNames []string

	toolNames = append(toolNames, WikipediaInit())
	toolNames = append(toolNames, DuckDuckGoInit())
	toolNames = append(toolNames, NowTimeToolInit())

	for _, toolName := range toolNames {
		fmt.Println("工具:", toolName, "初始化!")
	}
}
