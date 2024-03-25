package conf

import (
	"encoding/json"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 这里不采用直接暴露变量的方式，直接使用函数
var config *Config

// 补充逻辑
func C() *Config {
	// 如果配置为空
	if config == nil {
		// 给个默认值
		config = DefaultConfig()
	}
	return config
}

// 程序配置对象，启动时会读取配置，为程序提供一个全局变量
// 把配置对象做成全局变量
type Config struct {
	MySQL *MySQL `json:"mysql" yaml:"mysql" toml:"mysql"`
}

// 默认值，使得程序没有配置默认可以运行
func DefaultConfig() *Config {
	return &Config{
		MySQL: &MySQL{
			Host:     "110.41.160.251",
			Port:     3310,
			DB:       "vblog",
			Username: "root",
			Password: "123456Aa.",
			Debug:    true,
		},
	}
}

// 自定义Stringger打印值
func (c *Config) String() string {
	jsondata, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("%p", c)
	}
	return string(jsondata)
}

// db对象也是一个单列模式
type MySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`

	// 判断这个私有属性, 来判断是否返回已有的对象
	db *gorm.DB
	l  sync.Mutex
}

// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (m *MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
	)
}

// 通过配置就能通过一个DB 实例
func (m *MySQL) GetDB() *gorm.DB {
	m.l.Lock()
	defer m.l.Unlock()

	if m.db == nil {
		db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.db = db

		// 补充Debug配置
		if m.Debug {
			m.db = db.Debug()
		}
	}

	return m.db
}

// 配置对象提供全局单列配置
func (c *Config) DB() *gorm.DB {
	return c.MySQL.GetDB()
}
