package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	Nacos struct {
		DataId      string
		Group       string
		Host        string
		Port        int
		NamespaceId string
	}
}

var Conf ViperConfig

func GetViperConfig() {
	viper.SetConfigFile("D:\\666zg4\\like-shop-three\\common\\config\\server.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("viper success")
}
