package conf

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/mundo-wang/wtool/wlog"
	"github.com/spf13/viper"
	"os"
)

var (
	Config YamlConfig
)

//go:embed config/*.yml
var configsFS embed.FS // 把所有配置文件打包进embed.FS对象

func InitConfig() error {
	env := os.Getenv("env")
	if env == "" {
		env = "dev"
	}
	configFile := fmt.Sprintf("config/config-%s.yml", env)
	data, err := configsFS.ReadFile(configFile)
	if err != nil {
		wlog.Error("call configsFS.ReadFile failed").Err(err).Field("configFile", configFile).Log()
		return err
	}
	viper.SetConfigType("yaml") // 指定配置文件类型这一步不可缺少
	err = viper.ReadConfig(bytes.NewReader(data))
	if err != nil {
		wlog.Error("call viper.ReadConfig failed").Err(err).Log()
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
