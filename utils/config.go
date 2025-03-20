package utils

import (
	"fmt"
	"github.com/mundo-wang/wtool/wlog"
	"github.com/spf13/viper"
	"im-chat/conf"
	"os"
)

var (
	Config conf.Config
)

func InitConfig() error {
	env := os.Getenv("env")
	if env == "" {
		env = "dev"
	}
	viper.SetConfigFile(fmt.Sprintf("conf/config-%s.yml", env))
	err := viper.ReadInConfig()
	if err != nil {
		wlog.Error("call viper.ReadInConfig failed").Err(err).Log()
		return err
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		wlog.Error("call viper.Unmarshal failed").Err(err).Log()
		return err
	}
	wlog.Info("InitConfig complete!").Log()
	return nil
}
