package tools

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/tool/wikipedia"
	"github.com/cloudwego/eino/components/tool"
)

var WikipediaTool tool.InvokableTool

func WikipediaInit() {

	ctx := context.Background()

	cfg := &wikipedia.Config{
		Language: "zh",
	}

	wikipediaTool, err := wikipedia.NewTool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	WikipediaTool = wikipediaTool
}
