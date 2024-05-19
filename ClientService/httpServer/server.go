package main

import (
	"client/client"
	_ "client/docs"
	"client/grpcServer/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/swaggo/http-swagger"
	"io"
	"log"
	"net/http"
)

// @title Client Microservice

type server struct {
	db repository.Database
}

func main() {
	db, err := repository.Connect(context.Background())
	if err != nil {
		fmt.Errorf("can't connect to database")
	}
	s := &server{db: db}

	http.HandleFunc("/orders/show", s.ShowOrders)
	http.HandleFunc("/orders/create", s.MakeOrder)

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ShowOrders godoc
// @Summary Show all orders
// @Description Get all orders
// @Tags orders
// @Produce json
// @Success 200
// @Router /orders/show [get]
func (s *server) ShowOrders(writer http.ResponseWriter, request *http.Request) {
	message, err := s.db.GetOrders(context.Background())
	if err != nil {
		log.Fatal("Can't get orders")
	}

	resp, err := json.MarshalIndent(message, "", " ")
	if err != nil {
		log.Fatal("Can't parse json")
	}

	_, err = writer.Write(resp)
	if err != nil {
		log.Fatal("Can't write")
	}
}

type CreateOrderRequest struct {
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
}

// MakeOrder godoc
// @Summary Create a new order
// @Description Make an order with the given product name and quantity
// @Tags orders
// @Accept  json
// @Produce json
// @Param order body CreateOrderRequest true "Create order request"
// @Success 200
// @Router /orders/create [post]
func (s *server) MakeOrder(writer http.ResponseWriter, request *http.Request) {
	c, err := client.NewClient()
	req := CreateOrderRequest{}

	body, _ := io.ReadAll(request.Body)

	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal("Can't read request body")
	}

	message, err := c.MakeOrder(req.Name, req.Quantity)
	if err != nil {
		log.Fatal("Can't create order")
	}

	resp, err := json.MarshalIndent(message, "", " ")
	if err != nil {
		log.Fatal("Can't parse json")
	}

	_, err = writer.Write(resp)
	if err != nil {
		log.Fatal("Can't write")
	}
}
