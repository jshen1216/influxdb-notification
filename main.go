package main

import (
	"influx-alert-api/global"
	"influx-alert-api/router"
	"influx-alert-api/utils"
)

// @title InfluxDB Alert Api
// @version 2.0
// @description Golang API
// @termsOfService http://swagger.io/terms/
// @contact.name Jessie Shen
// @contact.email jessie.shen@bimap.co
// @host localhost:5555
// @query.collection.format multi
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @schemes http
func main() {

	utils.LoadEnvironment()
	r := router.LoadRouter()
	r.Run(global.EnvConfig.Server.Port)
}
