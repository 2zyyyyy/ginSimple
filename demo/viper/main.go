package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

// Viper是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式

type Gin struct {
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

type MySql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type YamlConfig struct {
	Gin   `yaml:"gin"`
	MySql `yaml:"mysql"`
}

var yamlConfig = new(YamlConfig)

func main() {
	vip := viper.New()
	// 设置配置文件名，没有后缀
	vip.SetConfigName("app")
	// 设置读取文件格式为: yaml
	vip.SetConfigType("yaml")
	// 设置配置文件目录(可以设置多个,优先级根据添加顺序来)
	vip.AddConfigPath("./config")
	// 读取解析
	if err := vip.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found, err:", err)
			return
		} else {
			fmt.Println("config file find success, but parse file error.", err)
			return
		}
	}

	// 映射到结构体
	if err := vip.Unmarshal(&yamlConfig); err != nil {
		fmt.Println("config file unmarshal struct failed, err:", err)
	}

	// 格式化输出
	formatJson(yamlConfig.MySql)
	formatJson(yamlConfig.Gin)
}

// struct转json
func formatJson(data interface{}) {
	byteArray, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("json marshal indent failed, err:", err)
	}
	fmt.Println(string(byteArray))
}
