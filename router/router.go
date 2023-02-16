package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "influx-alert-api/controller"

	_ "influx-alert-api/docs"
	"influx-alert-api/global"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoadRouter() *gin.Engine {

	gin.SetMode(global.EnvConfig.Server.Mode)
	router := gin.Default()

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	apiv2 := router.Group("api")
	apiv2.POST("Alert/Line", AlertLine)
	apiv2.POST("Alert/Email", AlertEmail)

	return router
}
