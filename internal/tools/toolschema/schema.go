package toolschema

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type NowTimeReq struct{}

type MdReaderReq struct {
	Path string `json:"path"`
}

type MdWriterReq struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}
