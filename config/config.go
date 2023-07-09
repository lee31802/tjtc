package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Configuration struct {
	App App `mapstructure:"app" json:"app" yaml:"app"`
}

type Application struct {
	ConfigViper *viper.Viper
	Config      Configuration
}

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}

func InitializeConfig() *viper.Viper {
	// 设置配置文件路径
	config := "config.yaml"
	// 生产环境可以通过设置环境变量来改变配置文件路径
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		logrus.Error("InitializeConfig err:%v", err)
	}

	// 监听配置文件
	v.WatchConfig()
	globalConig := Application{}
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&globalConig); err != nil {
			logrus.Error("InitializeConfig err:%v", err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&globalConig); err != nil {
		logrus.Error("InitializeConfig err:%v", err)
	}

	return v
}
