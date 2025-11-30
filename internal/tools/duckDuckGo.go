package tools

import (
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo/v2"
)

// duck搜索工具配置与初始化

func DuckDuckGoInit() string {

	ctx := context.Background()

	cfg := &duckduckgo.Config{
		ToolDesc:   "search for information by duckduckgo,and get url",
		MaxResults: 3,
		Region:     duckduckgo.RegionWT, //地区,没中国配置混蛋
		ToolName:   "duckduckgo_text_search",
	}

	duckduckgo, err := duckduckgo.NewTextSearchTool(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	duckgoInfo, err := duckduckgo.Info(ctx)
	if err != nil {
		log.Fatal("duckgosearch注册:", err)
	}

	Register(duckgoInfo, duckduckgo)

	return duckgoInfo.Name
}
