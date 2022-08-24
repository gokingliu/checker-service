package configs

import "git.code.oa.com/trpc-go/trpc-go/errs"

// Status code msg 返回状态
type Status struct {
	Code int32
	Msg  string
}

var (
	ResOk                        = Status{0, "Success"}
	ResFail                      = Status{-1, ""}
	InnerLoadYamlError           = Status{-100, "内部读取Yaml失败"}
	InnerUnmarshalYamlError      = Status{-101, "内部解析Yaml失败"}
	ClientPostParamsRequestError = Status{100, "The params of POST method only support map[string]interface{} type!"}
	ClientGetQueryRequestError   = Status{101, "The params of GET method only support string type!"}
)

// New 错误构造方法
func New(err Status) error {
	return errs.New(int(err.Code), err.Msg)
}
