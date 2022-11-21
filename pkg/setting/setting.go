package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
	"github.com/spf13/viper"
)

var (
	CFG          *viper.Viper
	Cfg          *ini.File
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

func init() {
	CFG = viper.New()
	CFG.AutomaticEnv()
	CFG.ReadInConfig()
	// var err error
	// Cfg, err = ini.Load("conf/config.ini")
	// if err != nil {
	// 	logging.Fatal("Config load failed err:", err.Error())
	// }
	loadBase()
	loadServer()
	loadAPP()

}

func loadBase() {
	log.Println("runmode:", RunMode)
	RunMode = CFG.GetString("RUN_MODE")
	log.Println("runmode:", RunMode)
	//RunMode = Cfg.Section("").Key("RUN_MODE").MustString("DEBUG")
}

func loadServer() {
	// sec, err := Cfg.GetSection("server")
	// if err != nil {
	// 	logging.Fatal("server config init failed :", err.Error())
	// }
	HttpPort = CFG.GetInt("HTTP_PORT")
	//HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	//ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	//WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	ReadTimeout = CFG.GetDuration("READ_TIMEOUT")
	WriteTimeout = CFG.GetDuration("WRITE_TIMEOUT")
}

func loadAPP() {
	// sec, err := Cfg.GetSection("app")
	// if err != nil {
	// 	logging.Fatal("app config init failed :", err.Error())
	// }
	PageSize = CFG.GetInt("PAGE_SIZE")
	JwtSecret = CFG.GetString("JWT_SECRET")
	//PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	//JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}
