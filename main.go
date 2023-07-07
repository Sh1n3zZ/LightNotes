package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	app := gin.Default()

	if viper.GetBool("debug") {
		gin.SetMode(gin.DebugMode)
		app.Use(gin.Logger())
		app.Use(gin.Recovery())
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := app.Run(fmt.Sprintf(":%s", viper.GetString("server.port"))); err != nil {
		panic(err)
	}
}
