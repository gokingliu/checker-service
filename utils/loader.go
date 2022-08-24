package utils

import (
	"checker/configs"
	"git.code.oa.com/trpc-go/trpc-go/config"
)

// LoadYaml 读取 yaml 配置信息
func LoadYaml(path string) (map[string]interface{}, error) {
	// 读取本地 yaml
	c, err := config.Load(path, config.WithCodec("yaml"), config.WithProvider("file"))
	// 读取 Yaml 出错
	if err != nil {
		return nil, configs.New(configs.InnerLoadYamlError)
	}
	// 定义 yaml content 类型
	yamlContent := make(map[string]interface{})
	// 解析 content 并赋值
	if err := c.Unmarshal(&yamlContent); err != nil {
		return nil, configs.New(configs.InnerUnmarshalYamlError)
	}

	return yamlContent, nil
}
