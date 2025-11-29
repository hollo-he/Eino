package tools

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo/v2"
	"github.com/cloudwego/eino/components/tool"
)

var DuckDuckGo tool.InvokableTool

func DuckDuckGoInit() {

	ctx := context.Background()

	cfg := &duckduckgo.Config{
		ToolDesc:   "search for information by duckduckgo,and get url",
		MaxResults: 3,
		Region:     duckduckgo.RegionWT,
	}

	duckduckgo, err := duckduckgo.NewTextSearchTool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	DuckDuckGo = duckduckgo
}
