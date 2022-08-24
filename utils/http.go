package utils

import (
	"checker/configs"
	"checker/models"
	"context"
	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/codec"
	"git.code.oa.com/trpc-go/trpc-go/http"
	reqHttp "net/http"
	"strings"
)

// RequestHandle 请求泛 HTTP 服务
func RequestHandle(httpName string, httpPath string, method string, headers map[string]string,
	cookies map[string]string, httpReq interface{}, httpRsp *models.HttpRsp) error {
	// 创建 ClientProxy, 设置协议为 HTTP 协议,序列化为 Json
	httpCli := http.NewClientProxy(httpName, client.WithProtocol("http"),
		client.WithSerializationType(codec.SerializationTypeJSON))
	reqHeader := &http.ClientReqHeader{}
	// 请求方法
	reqHeader.Method = method
	// 请求头
	if len(headers) != 0 {
		for key, value := range headers {
			reqHeader.AddHeader(key, value)
		}
	}
	// 设置 Cookie
	cookiesSlice := make([]string, 0)
	if len(cookies) != 0 {
		for name, value := range cookies {
			cookie := &reqHttp.Cookie{Name: name, Value: value, HttpOnly: false}
			cookiesSlice = append(cookiesSlice, cookie.String())
		}
	}
	reqHeader.AddHeader("Cookie", strings.Join(cookiesSlice, ""))
	// 返回头
	rspHead := &http.ClientRspHeader{}
	if method == "POST" {
		// 请求体
		mapReq, ok := httpReq.(map[string]interface{})
		// 发送 HTTP POST 请求
		if ok {
			err := httpCli.Post(
				context.Background(),
				httpPath,
				mapReq,
				&httpRsp,
				client.WithReqHead(reqHeader),
				client.WithRspHead(rspHead))
			if err != nil {
				return err
			}
		} else {
			return configs.New(configs.ClientPostParamsRequestError)
		}
	}
	if method == "GET" {
		// Query 参数
		query, ok := httpReq.(string)
		// 发送 HTTP GET 请求
		if ok {
			httpPath += query
			err := httpCli.Get(
				context.Background(),
				httpPath,
				&httpRsp,
				client.WithReqHead(reqHeader),
				client.WithRspHead(rspHead))
			if err != nil {
				return err
			}
		} else {
			return configs.New(configs.ClientGetQueryRequestError)
		}
	}

	return nil
}
