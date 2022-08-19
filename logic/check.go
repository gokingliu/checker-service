package logic

import (
	"checker/utils"
)

// GetHealthLogic 获取进程/文件是否存在，返回检查结果、不存在进程/文件、错误信息
func GetHealthLogic(key string) (bool, []string, error) {
	// 读取 Yaml 配置
	yamlList, err := utils.LoadYaml(key)
	if err != nil {
		return false, nil, err
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

	return flag, list, nil
}
