package tools

import (
	"Eino/internal/tools/toolschema"
	"context"
	"log"
	"os"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

func writerFile(ctx context.Context, req *toolschema.MdWriterReq) (*toolschema.Result, error) {
	err := os.WriteFile(req.Path, []byte(req.Content), 0644)
	if err != nil {
		return nil, err
	}

	return &toolschema.Result{
		Code: 200,
		Msg:  "OK",
	}, nil
}
func NewMdWriterTool() tool.InvokableTool {
	params := schema.NewParamsOneOfByParams(
		map[string]*schema.ParameterInfo{
			"path": &schema.ParameterInfo{
				Type:     schema.String,
				Desc:     "File path that needs to be modified",
				Required: true,
			},
			"content": &schema.ParameterInfo{
				Type: schema.String,
				Desc: "File content",
			},
		})
	return utils.NewTool(
		&schema.ToolInfo{
			Name:        "fileWriter",
			Desc:        "Update Markdown file",
			ParamsOneOf: params,
		}, writerFile)
}
func MdWriterInit() string {
	ctx := context.Background()
	mdWriter := NewMdWriterTool()
	mdWriterInfo, err := mdWriter.Info(ctx)
	if err != nil {
		log.Fatal(err)
	}
	Register(mdWriterInfo, mdWriter)
	return mdWriterInfo.Name
}
