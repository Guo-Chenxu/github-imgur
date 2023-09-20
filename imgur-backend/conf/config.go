package conf

import "gopkg.in/ini.v1"

// AppConfig 应用程序配置
type AppConfig struct {
	HttpPort      int `ini:"HttpPort"`
	*GithubConfig `ini:"github"`
	*CDNConfig    `ini:"cdn"`
}

// GithubConfig github配置
type GithubConfig struct {
	Accept        string `ini:"Accept"`
	Authorization string `ini:"Authorization"`
	User          string `ini:"User"`
	Path          string `ini:"Path"`
	Message       string `ini:"Message"`
}

// CDN cdn加速
type CDNConfig struct {
	CDN string `ini:"CDN"`
}

var Conf = new(AppConfig)

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
