package cmd

import (
	"github.com/RabbitMQ-Example/config"
	"github.com/RabbitMQ-Example/receiver"
	"github.com/spf13/cobra"
	"log"
)

// serveCmd represents the serve command
var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive...",
	Long:  `Receiving...`,
	Run:   receive,
}

func init() {
	rootCmd.AddCommand(receiveCmd)
}

func receive(cmd *cobra.Command, args []string) {
	ch, mq, err := config.GetMQConnection()
	if err != nil {
		log.Println(err)
	}
	receiver.Receive(ch, mq)

}
