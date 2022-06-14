package tools

/*
	配置文件快速读取帮助方法
*/

import (
	"errors"
	"gopkg.in/yaml.v2"
)

func ConfOfByte(dataByte []byte, data any) error {
	if err := yaml.Unmarshal(dataByte, data); err != nil {
		return errors.New(Fmt(`config.yml - %s - err: %s`, "yaml 格式化失败", err))
	}
	return nil
}
