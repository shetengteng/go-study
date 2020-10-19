package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// viper 使用 https://github.com/mitchellh/mapstructure 将配置反序列化到Config对象中 不写默认通过名称的tag查找

type Config struct {
	Port    int
	Version string `mapstructure:"version"`
	Mysql   struct { // 使用匿名结构体
		Port     int
		Url      string
		Username string `mapstructure:"user_name"` // 指定读取的key
	}
}

var Conf = new(Config)

// 读取配置文件，注意读取yaml文件的时候大小写不敏感，会被相同key覆盖（key不区分大小写，但是yaml对大小写敏感）
func main() {

	//viper.SetConfigFile("./conf2/a.yaml") // 指定文件路径

	viper.SetConfigName("config") // 指定文件名称
	viper.SetConfigType("yaml")   // 指定文件类型
	viper.AddConfigPath(".")      // 在工作路径中先查找
	viper.AddConfigPath("./conf") // 添加搜索路径，可以添加多个 先读取到相同的文件，后期不会继续读取
	//viper.AddConfigPath("./conf2") // 添加搜索路径，可以添加多个

	err := viper.ReadInConfig() // 读取配置文件
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
	fmt.Println(Conf.Mysql.Port)

}
