package main

import (
	"fmt"

	"github.com/cqroot/garden/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func setDefaultConfig() {
	viper.SetDefault("bind_ip", "127.0.0.1")
	viper.SetDefault("bind_port", "6003")
}

func main() {
	setDefaultConfig()
	common.InitConfig("object-server")
	common.InitLogger()

	r := gin.Default()

	v1Group := r.Group("/v1")
	{
		v1Group.HEAD("/:account/:bucket/:object", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"account": c.Param("account"),
				"bucket":  c.Param("bucket"),
				"object":  c.Param("object"),
			})
		})
		v1Group.PUT("/:account/:bucket/:object", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"account": c.Param("account"),
				"bucket":  c.Param("bucket"),
				"object":  c.Param("object"),
			})
		})
		v1Group.GET("/:account/:bucket/:object", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"account": c.Param("account"),
				"bucket":  c.Param("bucket"),
				"object":  c.Param("object"),
			})
		})
		v1Group.DELETE("/:account/:bucket/:object", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"account": c.Param("account"),
				"bucket":  c.Param("bucket"),
				"object":  c.Param("object"),
			})
		})
	}

	r.Run(fmt.Sprintf("%s:%s", viper.GetString("bind_ip"), viper.GetString("bind_port")))
}
