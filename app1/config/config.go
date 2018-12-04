package config

import "os"

const (
	ENV_STR = "env"
	ENV_DEV = "dev"
	ENV_RELEASE = "release"
)

type Config struct {
	MysqlDb MysqlInfo
}

type MysqlInfo struct {
	Db   string
	User string
	Pwd  string
	Addr string
}


var ConfigInfo *Config

func LoadConfig() (*Config, error) {
	// load config do something
	env := os.Getenv(ENV_STR)
	if env == ENV_DEV {
		ConfigInfo = &Config{MysqlDb: MysqlInfo{Db: "testdb", User: "testuser", Pwd: "abcdefa12", Addr: "192.168.0.15"}}
	} else {
		ConfigInfo = &Config{MysqlDb: MysqlInfo{Db: "realdb", User: "user", Pwd: "abcdefa12", Addr: "192.168.0.5"}}
	}
	return ConfigInfo, nil
}