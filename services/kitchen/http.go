package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gokuls-codes/go-grpc/services/common/genproto/orders"
	"github.com/gokuls-codes/go-grpc/services/common/genproto/orders/utils"
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

	conn, err := NewGRPCClient(":8000")

	if err != nil {
		log.Println("Error connecting to gRPC server:", err)
		return err
	}

	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)
		ctx, cancel := context.WithTimeout(r.Context(), time.Second * 2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerId: 24,
			ProductId:  1,
			Quantity:   2,
		})

		if err != nil {
			log.Println(err)
			utils.WriteError(w, http.StatusInternalServerError, err)
		}

		res, err := c.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerId: 2,
		})

		if err != nil {
			log.Println(err)
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))
		err = t.Execute(w, res.Orders)
		if err != nil {
			log.Println(err)
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
	})
	log.Println("Starting http server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<title>Kitchen Orders</title>
	</head>

<body>
	<h1>Kitchen Orders List</h1>
	<table border="1">

		<tr>
			<th>Order ID</th>
			<th>Customer ID</th>
			<th>Product ID</th>
			<th>Quantity</th>
		</tr>

		{{range .}}
		<tr>
			<td>{{.OrderID}}</td>
			<td>{{.CustomerID}}</td>
			<td>{{.ProductID}}</td>
			<td>{{.Quantity}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>`