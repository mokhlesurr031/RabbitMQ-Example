package main

import (
	"github.com/RabbitMQ-Example/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Println(err)
	}
}
