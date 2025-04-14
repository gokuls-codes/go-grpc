package services

import (
	"context"

	"github.com/gokuls-codes/go-grpc/services/common/genproto/orders"
)

var ordersList = make([]*orders.Order, 0)

type OrderSerice struct {


}

func NewOrderService() *OrderSerice {
	return &OrderSerice{}
}

func (s *OrderSerice) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersList = append(ordersList, order)
	return nil
}