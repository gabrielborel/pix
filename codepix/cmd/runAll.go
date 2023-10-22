package cmd

import (
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gabrielborel/pix/codepix/application/grpc"
	"github.com/gabrielborel/pix/codepix/application/kafka"
	"github.com/gabrielborel/pix/codepix/infra/db"
	"github.com/spf13/cobra"
)

var (
	gRPCPortNumber int
)

var runAllCmd = &cobra.Command{
	Use:   "runAll",
	Short: "Command used to start all services (kafka and grpc)",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		go grpc.StartGrpcServer(database, gRPCPortNumber)

		producer, err := kafka.NewKafkaProducer()
		if err != nil {
			panic(err)
		}

		deliveryChan := make(chan ckafka.Event)
		go kafka.DeliveryReport(deliveryChan)
		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(runAllCmd)
	runAllCmd.Flags().IntVarP(&gRPCPortNumber, "port", "p", 50051, "gRPC port")
}
