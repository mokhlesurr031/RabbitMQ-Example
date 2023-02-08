package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// cfgFile store the configuration file name
	//cfgFile                 string
	//verbose, prettyPrintLog bool
	rootCmd = &cobra.Command{
		Use:   "mq",
		Short: "RabbitMQ on test",
		Long:  `RabbitMQ on test`,
	}
)

//func init() {
//	cobra.OnInitialize(initConfig)
//}
//
//func initConfig() {
//	log.Println("Loading configurations")
//	log.Println("Configurations loaded successfully!")
//}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
