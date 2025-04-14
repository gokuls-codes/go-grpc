package main

import (
	"log"
	"net"

	handlers "github.com/gokuls-codes/go-grpc/services/orders/handlers/orders"
	"github.com/gokuls-codes/go-grpc/services/orders/services"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *gRPCServer) Run() error {

	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderService := services.NewOrderService()
	handlers.NewGrpcOrdersService(grpcServer, orderService)


	// Register your gRPC services here

	log.Println("Starting gPRC server on", s.addr)
	return grpcServer.Serve(lis)
}