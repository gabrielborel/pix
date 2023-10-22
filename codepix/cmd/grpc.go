package cmd

import (
	"fmt"
	"os"

	"github.com/gabrielborel/pix/codepix/application/grpc"
	"github.com/gabrielborel/pix/codepix/infra/db"
	"github.com/spf13/cobra"
)

var portNumber int
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Command used to start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		fmt.Println(portNumber)
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC port")
}
