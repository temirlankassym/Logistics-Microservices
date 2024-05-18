package main

import (
	"context"
	"encoding/json"
	_ "final/SupplierService/docs"
	"final/SupplierService/grpcServer/repository"
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
	http.HandleFunc("/suppliers/show", s.ShowSuppliers)
	http.HandleFunc("/suppliers/add", s.AddSupplier)
	http.HandleFunc("/suppliers/delete", s.DeleteSupplier)
	http.HandleFunc("/deliveries/show", s.ShowDeliveries)

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8082", nil))
}

func (s *server) ShowSuppliers(writer http.ResponseWriter, request *http.Request) {
	stock, err := s.db.ShowSuppliers(context.Background())

	resp, err := json.MarshalIndent(stock, "", " ")
	if err != nil {
		log.Fatal("Can't parse json")
	}

	_, err = writer.Write(resp)
	if err != nil {
		log.Fatal("Can't write")
	}
}

type AddSupplierRequest struct {
	ProductName string `json:"name"`
	CompanyName string `json:"company_name"`
	DaysToShip  int32  `json:"days_to_ship"`
}

func (s *server) AddSupplier(writer http.ResponseWriter, request *http.Request) {
	req := AddSupplierRequest{}

	body, _ := io.ReadAll(request.Body)

	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal("Can't read request body")
	}

	// check if request is valid
	if req.ProductName == "" || req.CompanyName == "" || req.DaysToShip == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.db.AddSupplier(context.Background(), req.ProductName, req.CompanyName, req.DaysToShip)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

type DeleteSupplierRequest struct {
	ProductName string `json:"name"`
}

func (s *server) DeleteSupplier(writer http.ResponseWriter, request *http.Request) {
	req := DeleteSupplierRequest{}

	body, _ := io.ReadAll(request.Body)

	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal("Can't read request body")
	}

	// check if request is valid
	if req.ProductName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.db.DeleteSupplier(context.Background(), req.ProductName)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}

func (s *server) ShowDeliveries(writer http.ResponseWriter, request *http.Request) {
	deliveries, err := s.db.ShowDeliveries(context.Background())

	resp, err := json.MarshalIndent(deliveries, "", " ")
	if err != nil {
		log.Fatal("Can't parse json")
	}

	_, err = writer.Write(resp)
	if err != nil {
		log.Fatal("Can't write")
	}
}
