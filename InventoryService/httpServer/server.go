package main

import (
	"context"
	"encoding/json"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "inventory/docs"
	"inventory/grpcServer/repository"
	"io"
	"log"
	"net/http"
)

// @title Inventory Microservice API
// @description API for managing inventory, including adding, deleting, and viewing stock.

type server struct {
	db repository.Database
}

func main() {
	db, err := repository.Connect(context.Background())
	if err != nil {
		fmt.Errorf("can't connect to database")
	}

	s := &server{db: db}
	http.HandleFunc("/stock/show", s.ShowStock)
	http.HandleFunc("/stock/add", s.AddProduct)
	http.HandleFunc("/stock/delete", s.DeleteProduct)

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

// ShowStock @Summary Show Stock
// @Description Get the list of all products in stock
// @Tags stock
// @Produce json
// @Success 200 {array} repository.Product
// @Failure 500 {object} map[string]string
// @Router /stock/show [get]
func (s *server) ShowStock(writer http.ResponseWriter, request *http.Request) {
	stock, err := s.db.ShowStock(context.Background())

	resp, err := json.MarshalIndent(stock, "", " ")
	if err != nil {
		log.Fatal("Can't parse json")
	}

	_, err = writer.Write(resp)
	if err != nil {
		log.Fatal("Can't write")
	}
}

// CreateProductRequest represents the request body for adding a product
type CreateProductRequest struct {
	Name        string `json:"name"`
	Quantity    int32  `json:"quantity"`
	Description string `json:"description"`
}

// AddProduct @Summary Add Product
// @Description Add a new product to the stock
// @Tags stock
// @Accept json
// @Produce json
// @Param product body CreateProductRequest true "Product to add"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/add [post]
func (s *server) AddProduct(writer http.ResponseWriter, request *http.Request) {
	req := CreateProductRequest{}

	body, _ := io.ReadAll(request.Body)

	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal("Can't read request body")
	}

	// check if req.Name and req.Quantity are empty
	if req.Name == "" || req.Quantity == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.db.AddProduct(context.Background(), req.Name, req.Quantity, req.Description)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

// DeleteProductRequest represents the request body for deleting a product
type DeleteProductRequest struct {
	Name string `json:"name"`
}

// DeleteProduct @Summary Delete Product
// @Description Delete a product from the stock
// @Tags stock
// @Accept json
// @Produce json
// @Param product body DeleteProductRequest true "Product to delete"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock/delete [delete]
func (s *server) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	req := DeleteProductRequest{}

	body, _ := io.ReadAll(request.Body)

	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal("Can't read request body")
	}

	// check if req.Name and req.Quantity are empty
	if req.Name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.db.DeleteProduct(context.Background(), req.Name)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}
