package cmd

import (
	"github.com/RabbitMQ-Example/config"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "mq",
		Short: "RabbitMQ on test",
		Long:  `RabbitMQ on test`,
	}
)

func init() {
	config.LoadRMQ()
}

func Execute() error {
	return rootCmd.Execute()
}
