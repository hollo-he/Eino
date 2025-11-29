package tools

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo/v2"
	"github.com/cloudwego/eino/components/tool"
)

// duck搜索工具配置与初始化
var DuckDuckGo tool.InvokableTool

func DuckDuckGoInit() {

	ctx := context.Background()

	cfg := &duckduckgo.Config{
		ToolDesc:   "search for information by duckduckgo,and get url",
		MaxResults: 3,
		Region:     duckduckgo.RegionWT, //地区,没中国配置混蛋
	}

	duckduckgo, err := duckduckgo.NewTextSearchTool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	DuckDuckGo = duckduckgo
}
