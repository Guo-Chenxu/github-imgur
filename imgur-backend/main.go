package main

import (
	"fmt"
	"imgur-backend/conf"
	"imgur-backend/router"
	"os"
)

const defaultConfFile = "./conf/config.ini"

func main() {
	confFile, bed := defaultConfFile, ""
	for i, arg := range os.Args {
		switch arg {
		case "-c":
			confFile = os.Args[i+1]
		case "-b":
			bed = os.Args[i+1]
		}
	}

	// 加载配置文件
	if err := conf.Init(confFile); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	if bed != "" {
		conf.Conf.Bed = bed
	}

	r := router.SetUpRouter()
	if err := r.Run(fmt.Sprintf(":%d", conf.Conf.HttpPort)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
