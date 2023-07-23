package main

import (
	"github.com/phuhao00/wddd/blog/internal/app/adapter"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote" // 必须导入，才能加载远程 key/value 配置
)

func main() {
	r := adapter.Router()
	err := viper.AddRemoteProvider("consul", "localhost:8500", "user/config") //
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("YAML")
	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
	r.Run(":8080")

}
