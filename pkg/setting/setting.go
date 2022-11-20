package setting

import (
	"time"

	"github.com/go-ini/ini"
	"github.com/jary-287/web-demo/pkg/logging"
)

var (
	Cfg          *ini.File
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/config.ini")
	if err != nil {
		logging.Fatal("Config load failed err:", err.Error())
	}
	loadBase()
	loadServer()
	loadAPP()

}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("DEBUG")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		logging.Fatal("server config init failed :", err.Error())
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadAPP() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		logging.Fatal("app config init failed :", err.Error())
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}
