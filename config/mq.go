package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type RMQ struct {
	Broker         string
	DefaultQueue   string
	TaskRetryDelay time.Duration
	TaskRetryCount int
}

var (
	mq RMQ
)

func MQ() *RMQ {
	return &mq
}

// LoadDB loads database configuration
func loadRMQ() {
	//viper.AddConfigPath("conf")
	viper.SetConfigFile("config.yml")
	er := viper.ReadInConfig()
	if er != nil {
		log.Println(er)
	}

	mq = RMQ{
		Broker:         viper.GetString("mq.Broker"),
		DefaultQueue:   viper.GetString("mq.DefaultQueue"),
		TaskRetryCount: viper.GetInt("mq.TaskRetryCount"),
		TaskRetryDelay: time.Duration(viper.GetInt("mq.TaskRetryDelay")),
	}
}
