package main

import (
	"log"
	"net/http"

	_ "github.com/jary-287/web-demo/docs"
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
	router := router.InitRouter()
	server := &http.Server{
		Handler:      router,
		ReadTimeout:  setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Panicln("the server start failed:", err.Error())
	}

}
