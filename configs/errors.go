package configs

import "git.code.oa.com/trpc-go/trpc-go/errs"

// Status code msg 返回状态
type Status struct {
	Code int32
	Msg  string
}

var (
	ResOk                   = Status{0, "Success"}
	InnerLoadYamlError      = Status{-100, "内部读取Yaml失败"}
	InnerUnmarshalYamlError = Status{-101, "内部解析Yaml失败"}
	ClientParamParsingError = Status{100, "客户端请求参数解析失败"}
)

// New 错误构造方法
func New(err Status) error {
	return errs.New(int(err.Code), err.Msg)
}
