package common

import (
	"github.com/spf13/viper"
)

func InitConfig(serverName string) {
	viper.SetDefault("log_level", "info")

	// viper.SetConfigName(serverName)
	// viper.SetConfigType("toml")
	// viper.AddConfigPath("/etc/simple-object-storage/")

	// err := viper.ReadInConfig()
	// if err != nil {
	//     panic(err)
	// }
}
