package cmd

import (
	"github.com/RabbitMQ-Example/config"
	"github.com/RabbitMQ-Example/sender"
	"github.com/spf13/cobra"
	"log"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Sender...",
	Long:  `Sending...`,
	Run:   send,
}

func init() {
	rootCmd.AddCommand(sendCmd)
}

func send(cmd *cobra.Command, args []string) {
	taskName := config.BasicTask
	taskBody := config.BasicTaskBody
	ch, mq, err := config.GetMQConnection(taskName)
	if err != nil {
		log.Println(err)
	}
	sender.Send(ch, mq, taskBody)

}
