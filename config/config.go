package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

var (
	appPath  = "/app"
	once     sync.Once
	instance *Config
)

func GetConfigInstance() *Config {
	once.Do(func() {
		viper.SetConfigName("env")
		if envPath := os.Getenv("CLUSTER_TOOL_PATH"); envPath != "" {
			appPath = envPath
		}
		viper.AddConfigPath(appPath)
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("read env file failed. ", err.Error())
		}
		instance = &Config{}
		viper.UnmarshalExact(instance)

	})

	return instance
}

type Config struct {
	UpstreamGRPC struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"upstream_grpc"`
}
