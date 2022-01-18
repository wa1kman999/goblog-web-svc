package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wa1kman999/goblog-web-svc/global"
)

// ConfigInit 初始化配置
func ConfigInit() error {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := v.Unmarshal(&global.GoBlogWebConfig); err != nil {
		return err
	}
	return nil
}
