package main

import (
	"imgur-backend/conf"
	"imgur-backend/router"
	"fmt"
	"os"
)

const defaultConfFile = "./conf/config.ini"

func main() {
	confFile := defaultConfFile
	if len(os.Args) > 2 {
		fmt.Println("use specified conf file: ", os.Args[1])
		confFile = os.Args[1]
	} else {
		fmt.Println("no configuration file was specified, use ./conf/config.ini")
	}
	// 加载配置文件
	if err := conf.Init(confFile); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	r := router.SetUpRouter()
	if err := r.Run(fmt.Sprintf(":%d", conf.Conf.HttpPort)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
