package config

import (
	"os"
	"strings"

	"github.com/sagungw/gotrunks/log"
	"github.com/spf13/viper"
)

func RedisAddress() string {
	return viper.GetString("redis.address")
}

func init() {
	viper.SetConfigName("config")

	if s := os.Getenv("PROJECT_DIR"); s != "" {
		viper.AddConfigPath(s)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("$PWD/")

	viper.SetDefault("redis.address", "redis://localhost:6379")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
	err := viper.MergeInConfig()
	if err != nil {
		log.From("config", "init").Infof("Error while reading config file %v", err)
	}
}
