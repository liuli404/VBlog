package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 该文件定义 配置对象的加载方法

// object <---> toml 配置文件
func LoadFromFile(filepath string) error {
	c := DefaultConfig()
	if _, err := toml.DecodeFile(filepath, c); err != nil {
		return err
	}
	config = c
	return nil
}

func LoadFromEnv() error {
	c := DefaultConfig()
	if err := env.Parse(c); err != nil {
		return err
	}
	config = c
	return nil
}
