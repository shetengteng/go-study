package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// viper 使用 https://github.com/mitchellh/mapstructure 将配置反序列化到Config对象中 不写默认通过名称的tag查找

type Config struct {
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

var Conf = new(Config)

// 读取配置文件
func main() {

	viper.SetConfigFile("./conf/config.yaml") // 指定文件路径
	err := viper.ReadInConfig()               // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	// 读取系统的变量
	err = viper.Unmarshal(Conf)
	if err != nil {
		panic(fmt.Errorf("Unmarshal config failed ,err %s \n", err))
	}

	fmt.Println(Conf.Port)
	fmt.Println(Conf.Version)

}
