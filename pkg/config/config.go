package config

import (
	"fmt"
	"os"
	"ship-manage-local/pkg/global"
	"ship-manage-local/pkg/logger"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var config *Conf

func Get() *Conf {
	return config
}

type Conf struct {
	Log    logger.Log `mapstructure:"log" yaml:"log"`
	Server Server     `mapstructure:"base_server" yaml:"base_server"`
}

type Server struct {
	Addr string `mapstructure:"addr" yaml:"addr"`
}

func New() (*Conf, error) {
	config = &Conf{
		Log: logger.NewLog(),
	}

	compareEnv()

	config.Log.InitLog()
	logrus.Info("配置信息加载成功!!!")

	return config, nil
}

func compareEnv() {
	// 日志信息
	config.Log.File = getEnvString("LOG_FILE", config.Log.File)
	config.Log.Level = getEnvString("LOG_LEVEL", config.Log.Level)
	config.Log.MaxSize = getEnvInt("LOG_SIZE", config.Log.MaxSize)
	config.Log.MaxAge = getEnvInt("LOG_AGE", config.Log.MaxAge)
	config.Log.MaxBackups = getEnvInt("LOG_BACKUPS", config.Log.MaxBackups)

	config.Server.Addr = getEnvString("ADDR", config.Server.Addr)
	if config.Server.Addr == "" {
		config.Server.Addr = global.DefaultAddr
	}
}

func (c *Conf) Show() {
	if b, err := yaml.Marshal(c); err != nil {
		return
	} else {
		fmt.Printf(`
-----------------------------------------------------------------------------------------
%s
-----------------------------------------------------------------------------------------
`, string(b))
	}
}

func getEnvString(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}

		return valueInt
	}

	return defaultValue
}
