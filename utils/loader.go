package utils

import (
	"checker/configs"
	"git.code.oa.com/trpc-go/trpc-go/config"
)

// LoadYaml 读取 yaml 配置信息
func LoadYaml(key string) ([]string, error) {
	// 读取本地 Yaml
	c, err := config.Load("conf/content.yaml", config.WithCodec("yaml"), config.WithProvider("file"))
	// 读取 Yaml 出错
	if err != nil {
		return nil, configs.New(configs.InnerLoadYamlError)
	}
	// 定义 Yaml content 类型
	yamlContent := make(map[string]map[string][]string)
	// 解析 content 并赋值
	if err := c.Unmarshal(&yamlContent); err != nil {
		return nil, configs.New(configs.InnerUnmarshalYamlError)
	}
	// 读取文件或进程
	list := yamlContent["content"][key]

	return list, nil
}
