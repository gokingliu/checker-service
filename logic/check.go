package logic

import (
	"checker/utils"
	"strings"
)

// GetHealthLogic 获取进程/文件是否存在，返回检查结果、不存在进程/文件、错误信息
func GetHealthLogic(key string) (bool, string, error) {
	// 读取 yaml 配置
	yamlList, err := utils.LoadYaml(key)
	if err != nil {
		return false, "", err
	}
	// 定义变量
	flag := true
	list := make([]string, 0)
	// 遍历进程/文件是否存在
	for i := 0; i < len(yamlList); i++ {
		if key == "process" {
			if !utils.CheckProcessExists(yamlList[i]) {
				flag = false
				list = append(list, yamlList[i])
			}
		} else {
			if !utils.CheckPathExists(yamlList[i]) {
				flag = false
				list = append(list, yamlList[i])
			}
		}
	}
	// 不存在的进程/文件
	str := ""
	if list != nil && len(list) > 0 {
		str = strings.Join(list, "|")
		if str != "" {
			if key == "process" {
				str = "缺失的进程：" + str + "；"
			} else {
				str = "缺失的文件：" + str
			}
		}
	}

	return flag, str, nil
}
