package main

import (
	"context"
	"encoding/json"
	_ "final/InventoryService/docs"
	"final/InventoryService/grpcServer/repository"
	"fmt"
	"github.com/swaggo/http-swagger"
	"io"
	"log"
	"net/http"
)

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

type CreateProductRequest struct {
	Name        string `json:"name"`
	Quantity    int32  `json:"quantity"`
	Description string `json:"description"`
}

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

type DeleteProductRequest struct {
	Name string `json:"name"`
}

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
