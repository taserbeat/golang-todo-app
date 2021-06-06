package setting

import (
	"log"

	"github.io/taserbeat/golang-todo-app/modules/utils"
	ini "gopkg.in/ini.v1"
)

type ConfigList struct {
	Port       string
	SQLDriver  string
	DbHost     string
	DbPort     string
	DbUser     string
	DbName     string
	DbPassword string
	LogFile    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:       cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver:  cfg.Section("db").Key("driver").String(),
		DbHost:     cfg.Section("db").Key("host").MustString("localhost"),
		DbPort:     cfg.Section("db").Key("port").MustString("15432"),
		DbName:     cfg.Section("db").Key("name").MustString("test_db"),
		DbUser:     cfg.Section("db").Key("user").MustString("root"),
		DbPassword: cfg.Section("db").Key("password").String(),
		LogFile:    cfg.Section("web").Key("logfile").String(),
	}
}
