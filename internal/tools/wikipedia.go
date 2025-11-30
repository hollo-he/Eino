package tools

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/tool/wikipedia"
)

// 配置维基百科

func WikipediaInit() string {

	ctx := context.Background()

	cfg := &wikipedia.Config{
		Language: "zh",
		ToolName: "wikipedia_search",
	}

	wikipediaTool, err := wikipedia.NewTool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	wikiinfo, err := wikipediaTool.Info(ctx)
	if err != nil {
		log.Fatal("维基百科注册:", err)
	}
	Register(wikiinfo, wikipediaTool)

	return wikiinfo.Name
}
