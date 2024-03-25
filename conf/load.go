package conf

import (
	"github.com/BurntSushi/toml"
)

// 该文件定义 配置对象的加载方法

// object <---> toml 配置文件
func LoadFromFile(filepath string) error {
	c := &Config{}
	if _, err := toml.DecodeFile(filepath, c); err != nil {
		return err
	}
	config = c
	return nil
}

func LoadFromEnv() error {
	// env.Parse()
	return nil
}
