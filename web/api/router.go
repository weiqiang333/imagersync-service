// author: weiqiang; date: 2022-03
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Default Url /
func Default(c *gin.Context) {
	configMapData := viper.GetStringMap("config")
	c.HTML(200, "default.tmpl", configMapData)
}
