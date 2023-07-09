package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"lightnotes/middleware"
)

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	app := gin.Default()

	if viper.GetBool("debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	{
		app.Use(middleware.BuiltinMiddleWare(ConnectMySQL(), ConnectRedis()))
		app.Use(middleware.ThrottleMiddleware())
		app.Use(AuthMiddleware())
	}
	{
		app.POST("/login", LoginAPI)
		app.POST("/anonymous/send", AnonymousSendAPI)
		app.GET("/anonymous/get", AnonymousGetAPI)
	}

	if err := app.Run(fmt.Sprintf(":%s", viper.GetString("server.port"))); err != nil {
		panic(err)
	}
}
