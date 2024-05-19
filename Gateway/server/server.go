package main

import (
	"encoding/json"
	"gateway/client"
	"io"
	"log"
	"net/http"
)

// @title Client Microservice API

func main() {
	http.HandleFunc("/orders/create", MakeOrder)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

// CreateOrderRequest represents the request body for creating an order
type CreateOrderRequest struct {
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
}

// MakeOrder godoc
// @Summary Create a new order
// @Description Make an order with the given product name and quantity
// @Tags orders
// @Accept json
// @Produce json
// @Param order body CreateOrderRequest true "Create order request"
// @Success 200 {object} map[string]string
// @Failure 400
// @Failure 500
// @Router /orders/create [post]
func MakeOrder(writer http.ResponseWriter, request *http.Request) {
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
