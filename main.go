package main

import (
	"context"
	generated "github.com/eyobmamo/go_grpc_implimantation-/grpc"
	"github.com/eyobmamo/go_grpc_implimantation-/handler"
	"github.com/eyobmamo/go_grpc_implimantation-/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpcSvc := service.NewServiceInstance(context.Background())

	handler := handler.NewGrpcHandler(grpcSvc)
	generated.RegisterConfigServiceServer(grpcServer, handler)

	log.Println("grpc server start lestining  on : 50051")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
