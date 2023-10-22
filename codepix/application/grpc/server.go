package grpc

import (
	"fmt"
	"log"
	"net"

	pb "github.com/gabrielborel/pix/codepix/application/grpc/pb"
	"github.com/gabrielborel/pix/codepix/application/usecase"
	"github.com/gabrielborel/pix/codepix/infra/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(db *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.NewPixKeyRepositoryDb(db)
	pixUseCase := usecase.NewPixKeyUseCase(pixRepository)
	pixGrpcService := NewPixGrpcService(*pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("could not listen to %s: %s", addr, err.Error())
	}

	log.Printf("gRPC server running at %s", addr)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("could not start gRPC server: %s", err.Error())
	}
}
