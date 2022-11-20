package main

import (
	"fmt"
	"syscall"

	"github.com/fvbock/endless"
	_ "github.com/jary-287/web-demo/docs"
	"github.com/jary-287/web-demo/pkg/logging"
	"github.com/jary-287/web-demo/pkg/setting"
	"github.com/jary-287/web-demo/router"
)

// @title           myweb API docs
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HttpPort)
	server := endless.NewServer(endPoint, router.InitRouter())
	server.BeforeBegin = func(add string) {
		logging.Info("Actual pid is", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		logging.Info("Server err:", err)
	}
	server.ListenAndServe()
}
