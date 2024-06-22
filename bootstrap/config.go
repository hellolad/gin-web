package bootstrap

import (
	"fmt"
	"gin-web/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func InitializeConfig() *viper.Viper {
	config := "config.yaml"
	if configEvn := os.Getenv("VIPER_CONFIG"); configEvn != "" {
		config = configEvn
	}
	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
	return v
}
