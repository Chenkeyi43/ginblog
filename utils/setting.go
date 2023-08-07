package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Zone        int
	AccessKey   string
	SecrtKey    string
	Bucket      string
	QiniuServer string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Println(cfg.Section("server").Key("AppMode"))
	fmt.Println(cfg.Section("server").KeyStrings())

	LoadServer(cfg)
	LoadDatabase(cfg)
	LoadQiniu(cfg)
}

func LoadServer(cfg *ini.File) {
	AppMode = cfg.Section("server").Key("AppMode").MustString("debug")
	HttpPort = cfg.Section("server").Key("HttpPort").MustString("debug")
	JwtKey = cfg.Section("server").Key("JwtKey").MustString("debug")
}

func LoadDatabase(cfg *ini.File) {
	Db = cfg.Section("database").Key("Db").MustString("debug")
	DbUser = cfg.Section("database").Key("DbUser").MustString("debug")

	DbHost = cfg.Section("database").Key("DbHost").MustString("debug")
	DbPort = cfg.Section("database").Key("DbPort").MustString("debug")
	DbPassWord = cfg.Section("database").Key("DbPassWord").MustString("debug")
	DbName = cfg.Section("database").Key("DbName").MustString("debug")
}

func LoadQiniu(cfg *ini.File) {

}
