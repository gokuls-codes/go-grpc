package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func NewGRPCClient (addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func main() {
	// Create a new HTTP server
	httpServer := NewHttpServer(":8080")

	// Run the HTTP server
	if err := httpServer.Run(); err != nil {
		log.Fatal(err)
	}
}