package cmd

import (
	"github.com/RabbitMQ-Example/config"
	"github.com/RabbitMQ-Example/receiver"
	"github.com/spf13/cobra"
	"log"
)

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
	taskName := config.BasicTask
	ch, mq, err := config.GetMQConnection(taskName)
	if err != nil {
		log.Println(err)
	}
	receiver.Receive(ch, mq)

}
