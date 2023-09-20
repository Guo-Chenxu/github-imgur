package conf

import "gopkg.in/ini.v1"

// AppConfig 应用程序配置
type AppConfig struct {
	HttpPort      int `ini:"HttpPort"`
	*GithubConfig `ini:"github"`
	*CDNConfig    `ini:"cdn"`
	*GiteeConfig  `ini:"gitee"`
	*BedConfig    `ini:"bed"`
}

// GithubConfig github配置
type GithubConfig struct {
	Accept        string `ini:"Accept"`
	Authorization string `ini:"Authorization"`
	Repo          string `ini:"Repo"`
	User          string `ini:"User"`
	Path          string `ini:"Path"`
	Message       string `ini:"Message"`
}

// CDN cdn加速
type CDNConfig struct {
	CDN string `ini:"CDN"`
}

// GiteeConfig gitee配置
type GiteeConfig struct {
	AccessToken string `ini:"AccessToken"`
	Owner       string `ini:"Owner"`
	Repo        string `ini:"Repo"`
	Path        string `ini:"Path"`
	Message     string `ini:"Message"`
	Branch       string `ini:"Branch"`
}

// Bed 图床选择
type BedConfig struct {
	Bed string `ini:"Bed"`
}

var Conf = new(AppConfig)

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
