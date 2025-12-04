package toolschema

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type NowTimeReq struct{}

type MdReaderReq struct {
	Path string `json:"path"`
}
