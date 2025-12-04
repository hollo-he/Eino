package tools

import (
	"context"
	"fmt"
	"log"
	"time"

	"Eino/internal/tools/toolschema"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

// 工具函数（无参数）
func nowTimeGet(ctx context.Context, req *toolschema.NowTimeReq) (*toolschema.Result, error) {
	nowtime := time.Now()
	return &toolschema.Result{
		Code: 200,
		Msg:  fmt.Sprintf("当前时间是 %s", nowtime.Format("2006-01-02 15:04:05")),
	}, nil
}

// 生成工具
func NewNowTimeTool() tool.InvokableTool {
	return utils.NewTool(
		&schema.ToolInfo{
			Name: "getNowTime",
			Desc: "获取当前系统时间（无参数）",
			ParamsOneOf: schema.NewParamsOneOfByParams(
				map[string]*schema.ParameterInfo{}, // 💡 无参数工具关键点
			),
		},
		nowTimeGet,
	)
}

func NowTimeToolInit() string {
	ctx := context.Background()
	nowTimeTool := NewNowTimeTool()
	nowTimeToolInfo, err := nowTimeTool.Info(ctx)
	if err != nil {
		log.Fatal(err)
	}
	Register(nowTimeToolInfo, nowTimeTool)
	return nowTimeToolInfo.Name
}
