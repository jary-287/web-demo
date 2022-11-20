package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jary-287/web-demo/middleware"
	"github.com/jary-287/web-demo/pkg/setting"
	"github.com/jary-287/web-demo/router/apis"
	v1 "github.com/jary-287/web-demo/router/apis/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", apis.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.JWT())

	api := r.Group("/api/v1")
	{
		//test
		api.GET("/health", v1.Ping)

		//tag
		api.GET("/tags", v1.GetTags)
		api.POST("/tags", v1.AddTag)
		api.PUT("/tags/:id", v1.EditTag)
		api.DELETE("/tags/:id", v1.DeleteTag)

		//blog
		api.GET("/articles", v1.GetArticles)
		api.POST("/articles", v1.AddArticle)
		api.GET("/articles/:id", v1.GetArticle)
		api.PUT("/articles/:id", v1.EditArticle)
		api.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
