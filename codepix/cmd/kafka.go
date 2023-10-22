package cmd

import (
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gabrielborel/pix/codepix/application/kafka"
	"github.com/gabrielborel/pix/codepix/infra/db"
	"github.com/spf13/cobra"
)

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Command used to start kafka server (start consuming transactions)",
	Run: func(cmd *cobra.Command, args []string) {
		producer, err := kafka.NewKafkaProducer()
		if err != nil {
			panic(err)
		}

		deliveryChan := make(chan ckafka.Event)
		database := db.ConnectDB(os.Getenv("env"))

		go kafka.DeliveryReport(deliveryChan)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
