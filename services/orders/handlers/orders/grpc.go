package handlers

import (
	"context"

	"github.com/gokuls-codes/go-grpc/services/common/genproto/orders"
	"github.com/gokuls-codes/go-grpc/services/orders/types"
	"google.golang.org/grpc"
)


type OrdersGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	grpcHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID: 42,
		CustomerID: 2,
		ProductID: 1,
		Quantity: 10,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}