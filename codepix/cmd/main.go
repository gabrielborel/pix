package main

import (
	"os"

	"github.com/gabrielborel/pix/codepix/application/grpc"
	"github.com/gabrielborel/pix/codepix/infra/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
