package handlers

import (
	"net/http"

	"github.com/gokuls-codes/go-grpc/services/common/genproto/orders"
	"github.com/gokuls-codes/go-grpc/services/common/genproto/orders/utils"
	"github.com/gokuls-codes/go-grpc/services/orders/types"
)


type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(ordersService types.OrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{
		ordersService: ordersService,
	}
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := utils.ParseJSON(r, &req)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order {
		OrderID: 23,
		CustomerID: req.GetCustomerId(),
		ProductID: req.GetProductId(),
		Quantity: req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	utils.WriteJSON(w, http.StatusCreated, res)
}