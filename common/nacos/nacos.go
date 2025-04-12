package nacos

import (
	"common/viper"
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type T struct {
	Mysql struct {
		User     string `json:"user"`
		Pass     string `json:"pass"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
	} `json:"mysql"`
	Redis struct {
		Addr string `json:"addr"`
		Pass string `json:"pass"`
		Db   int    `json:"db"`
	} `json:"redis"`
}

var Config T

func NewNacos() {
	con := viper.Conf
	clientConfig := constant.ClientConfig{
		NamespaceId:         con.Nacos.NamespaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      con.Nacos.Host,
			ContextPath: "/nacos",
			Port:        uint64(con.Nacos.Port),
			Scheme:      "http",
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: con.Nacos.DataId,
		Group:  con.Nacos.Group})
	if err != nil {
		fmt.Println("nacos连接失败")
		return
	}
	fmt.Println("nacos连接成功")
	err = json.Unmarshal([]byte(content), &Config)
	if err != nil {
		return
	}
	fmt.Println(content)
}
