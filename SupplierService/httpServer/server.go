package main

import (
	"context"
	"encoding/json"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"io"
	"log"
	"net/http"
	_ "supplier/docs"
	"supplier/grpcServer/repository"
)

// @title Supplier Microservice API
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

// ShowSuppliers @Summary Show Suppliers
// @Description Get the list of all suppliers
// @Tags suppliers
// @Produce json
// @Success 200 {array} repository.Supplier
// @Failure 500
// @Router /suppliers/show [get]
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

// AddSupplierRequest represents the request body for adding a supplier
type AddSupplierRequest struct {
	ProductName string `json:"name"`
	CompanyName string `json:"company_name"`
	DaysToShip  int32  `json:"days_to_ship"`
}

// AddSupplier @Summary Add Supplier
// @Description Add a new supplier to the system
// @Tags suppliers
// @Accept json
// @Produce json
// @Param supplier body AddSupplierRequest true "Supplier to add"
// @Success 200 {object} map[string]string
// @Failure 400
// @Failure 500
// @Router /suppliers/add [post]
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

// DeleteSupplierRequest represents the request body for deleting a supplier
type DeleteSupplierRequest struct {
	ProductName string `json:"name"`
}

// DeleteSupplier @Summary Delete Supplier
// @Description Delete a supplier from the system
// @Tags suppliers
// @Accept json
// @Produce json
// @Param supplier body DeleteSupplierRequest true "Supplier to delete"
// @Success 200 {object} map[string]string
// @Failure 400
// @Failure 500
// @Router /suppliers/delete [delete]
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

// ShowDeliveries @Summary Show Deliveries
// @Description Get the list of all deliveries
// @Tags deliveries
// @Produce json
// @Success 200 {array} repository.Delivery
// @Failure 400
// @Failure 500
// @Router /deliveries/show [get]
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
