package main

import (
	"log"
	"net/http"

	handlers "github.com/gokuls-codes/go-grpc/services/orders/handlers/orders"
	"github.com/gokuls-codes/go-grpc/services/orders/services"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := services.NewOrderService()
	ordersHandler := handlers.NewHttpOrdersHandler(orderService)

	ordersHandler.RegisterRouter(router)
	log.Println("Starting http server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}