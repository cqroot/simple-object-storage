package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/cqroot/simple-object-storage/internal/common"
	controllers "github.com/cqroot/simple-object-storage/internal/controllers/object"
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
		v1Group.PUT("/:account/:bucket/:object", controllers.PutObject)
		v1Group.GET("/:account/:bucket/:object", controllers.GetObject)
		v1Group.DELETE("/:account/:bucket/:object", controllers.DeleteObject)
	}

	r.Run(fmt.Sprintf("%s:%s", viper.GetString("bind_ip"), viper.GetString("bind_port")))
}
