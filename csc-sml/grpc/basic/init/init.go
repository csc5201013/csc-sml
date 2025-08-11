package init

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grpc/basic/global"
	"log"
)

func init() {
	InitConfig()
	InitMysql()
	InitNacos()

}

func InitConfig() {

	viper.SetConfigFile("./dev.yaml")
	viper.ReadInConfig()
	err := viper.Unmarshal(&global.AppConf)
	if err != nil {
		panic("动铁读取配置失败")
	}
	log.Println("动态读取配置成功", global.AppConf)

}

func InitMysql() {
	var err error
	conf := global.AppConf.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Post, conf.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		panic("数据库链接失败")

	}
	log.Println("数据库链接成功")

}

func InitNacos() {

	clientConfig := constant.ClientConfig{
		NamespaceId:         "77c9da26-98e6-4aac-9468-5a0e93d0362d", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "124.223.211.250",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	configClient, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-srv-json",
		Group:  "DEFAULT_GROUP",
	})

	log.Println("Nacos读取成功", content)

	//nacosConf := global.AppConf.Nacos
	//
	//clientConfig := constant.ClientConfig{
	//	NamespaceId:         nacosConf.Namespace,
	//	TimeoutMs:           5000,
	//	NotLoadCacheAtStart: true,
	//	LogDir:              "/tmp/nacos/log",
	//	CacheDir:            "/tmp/nacos/cache",
	//	LogLevel:            "debug",
	//}
	//
	//serverConfigs := []constant.ServerConfig{
	//	{
	//		IpAddr:      nacosConf.ServerAddr,
	//		ContextPath: "/nacos",
	//		Port:        8848,
	//		Scheme:      "http",
	//	},
	//}
	//configClient, err := clients.NewConfigClient(vo.NacosClientParam{
	//	ClientConfig:  &clientConfig,
	//	ServerConfigs: serverConfigs,
	//})
	//if err != nil {
	//	panic("Nacos客户端初始化失败: " + err.Error())
	//}
	//
	//// 修复配置获取参数格式
	//content, err := configClient.GetConfig(vo.ConfigParam{
	//	DataId: nacosConf.DataId,
	//	Group:  nacosConf.Group,
	//})
	//if err != nil {
	//	panic("nacos配置获取失败: " + err.Error())
	//}
	//log.Println("nacos配置获取成功", content)
}
