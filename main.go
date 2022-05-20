// author: weiqiang; date: 2022-05
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"imagersync-service/web/api"
	"imagersync-service/web/api/v1/imagersync"
)

func init() {
	pflag.String("configFile", string("configs/config.yaml"), "go config file")
	pflag.ErrHelp.Error()
}

func main() {
	loadConfig()

	router := engine()
	err := router.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(fmt.Errorf("web server faile: %s", err.Error()))
	}
}

// gin web run engine
func engine() *gin.Engine {
	router := gin.Default()
	router.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "./web/static")

	router.GET("/", api.Default)

	userAuth := loadAuthUsers()
	authorized := router.Group("/v1/image/rsync", gin.BasicAuth(gin.Accounts(userAuth)))
	authorized.GET("/", imagersync.GetImageRsync)
	authorized.POST("/", imagersync.PostImageRsync)

	return router
}

// load config and flag config
func loadConfig() {
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Println(err.Error())
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(viper.GetString("configFile"))
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func loadAuthUsers() map[string]string {
	return viper.GetStringMapString("rsyncserver.auth")
}
