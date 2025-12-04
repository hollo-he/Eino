package tools

import (
	"Eino/internal/tools/toolschema"
	"context"
	"errors"
	"log"
	"os"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

func mdReadder(ctx context.Context, req *toolschema.MdReaderReq) (*toolschema.Result, error) {
	data, err := os.ReadFile(req.Path)
	if err != nil {
		return nil, errors.New("读取路径:" + req.Path + ", 出现问题:" + err.Error())
	}

	return &toolschema.Result{
		Code: 200,
		Msg:  req.Path + "的内容为:\n" + string(data),
	}, nil
}

func NewMdReaderTool() tool.InvokableTool {
	params := schema.NewParamsOneOfByParams(
		map[string]*schema.ParameterInfo{
			"path": &schema.ParameterInfo{
				Type:     schema.String,
				Desc:     "File path to read",
				Required: true,
			},
		})
	return utils.NewTool(
		&schema.ToolInfo{
			Name:        "mdReader",
			Desc:        "Read a Markdown file and return its content",
			ParamsOneOf: params,
		}, mdReadder)
}

func MdReaderToolInit() string {

	ctx := context.Background()
	mdReadder := NewMdReaderTool()
	mdReadderInfo, err := mdReadder.Info(ctx)
	if err != nil {
		log.Fatal(err)
	}
	Register(mdReadderInfo, mdReadder)
	return mdReadderInfo.Name
}
