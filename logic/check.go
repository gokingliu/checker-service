package logic

import (
	"checker/configs"
	"checker/models"
	"checker/utils"
	"strings"
	"sync"
)

// HealthLogic 获取进程/文件是否存在，返回检查结果、不存在进程/文件、错误信息
func HealthLogic(key string) (bool, string, error) {
	// 读取 yaml 配置
	yamlContent, err := utils.LoadYaml("conf/content.yaml")
	// 读取 yaml 具体内容
	yamlList, ok := yamlContent["content"].(map[string]interface{})[key].([]interface{})
	if err != nil || !ok {
		return false, "", configs.New(configs.InnerUnmarshalYamlError)
	}
	// 定义变量
	flag := true
	list := make([]string, 0)
	// 遍历进程/文件是否存在
	for i := 0; i < len(yamlList); i++ {
		yamlItem, okItem := yamlList[i].(string)
		if !okItem {
			return false, "", configs.New(configs.InnerUnmarshalYamlError)
		}
		if key == "process" {
			if !utils.CheckProcessExists(yamlItem) {
				flag = false
				list = append(list, yamlItem)
			}
		} else {
			if !utils.CheckPathExists(yamlItem) {
				flag = false
				list = append(list, yamlItem)
			}
		}
	}
	// 不存在的进程/文件
	str := ""
	if list != nil && len(list) > 0 {
		str = strings.Join(list, "")
		if str != "" {
			if key == "process" {
				str = "缺失的进程: " + str
			} else {
				str = "缺失的文件: " + str
			}
		}
	}

	return flag, str, nil
}

// GetHealthLogic 批量获取多台机器的存活情况
func GetHealthLogic(reqType uint32) ([]string, error) {
	// 读取 yaml 配置
	yamlContent, err := utils.LoadYaml("conf/trpc_go.yaml")
	// 读取 yaml 具体内容
	yamlList, ok := yamlContent["client"].(map[string]interface{})["service"].([]interface{})
	if err != nil || !ok {
		return nil, configs.New(configs.InnerUnmarshalYamlError)
	}
	// client name 切片
	httpNameSlice := make([]string, 0)
	// 响应体
	var httpRsp models.HttpRsp
	// 响应体切片
	httpRspSlice := make([]models.HttpRsp, 0)
	// 客户端响应错误切片
	errSlice := make([]interface{}, 0)
	// client name 和 ip 对应切片
	httpIPNameMap := make([]string, 0)
	// 遍历 service 内容
	for i := 0; i < len(yamlList); i++ {
		clientName, okName := yamlList[i].(map[string]interface{})["name"].(string)
		clientTarget, okTarget := yamlList[i].(map[string]interface{})["target"].(string)
		// 解析出错
		if !okName || !okTarget {
			return nil, configs.New(configs.InnerUnmarshalYamlError)
		}
		// 判断客户端名称是否包含 Health
		if strings.Contains(clientName, "Health") {
			httpNameSlice = append(httpNameSlice, clientName)
			httpRspSlice = append(httpRspSlice, httpRsp)
			errSlice = append(errSlice, nil)
			httpIPNameMap = append(httpIPNameMap, clientTarget)
		}
	}
	// 请求头
	headers := map[string]string{}
	// 请求 cookie
	cookies := map[string]string{}
	// 请求参数
	httpReq := map[string]interface{}{"type": reqType}
	// go sync 协程异步等待
	nWait := sync.WaitGroup{}
	// 发起客户端请求
	requestHandle := func(index int) {
		err := utils.RequestHandle(
			httpNameSlice[index],
			"/trpc.checker.checkHealth.Check/Health",
			"POST",
			headers,
			cookies,
			httpReq,
			&httpRspSlice[index],
		)
		errSlice[index] = err
		nWait.Done()
	}
	// 遍历所有客户端请求
	for index := 0; index < len(httpNameSlice); index++ {
		nWait.Add(1)
		requestHandle(index)
	}
	nWait.Wait()
	// 错误信息切片
	msgErr := make([]string, 0)
	// 判断客户端请求是否出错
	for index, value := range errSlice {
		if value != nil {
			msgErr = append(msgErr, httpIPNameMap[index]+": 请求失败")
		} else if httpRspSlice[index].Code != 0 {
			msgErr = append(msgErr, httpIPNameMap[index]+": "+httpRspSlice[index].Msg)
		}
	}

	return msgErr, nil
}
