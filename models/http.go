package models

type HttpRsp struct {
	Code   int32
	Msg    string
	Result interface{}
}
