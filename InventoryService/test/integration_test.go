package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type Product struct {
	Id          int    `json:"Id"`
	Name        string `json:"ProductName"`
	Quantity    int32  `json:"Quantity"`
	Description string `json:"Description"`
}

func TestShowStock(t *testing.T) {
	var got []Product

	got = getProducts()

	expected := []Product{
		{
			Id:          1,
			Name:        "iPhone",
			Quantity:    0,
			Description: "Smartphone designed by Apple Inc.",
		},
		{
			Id:          2,
			Name:        "PlayStation",
			Quantity:    0,
			Description: "Gaming console developed by Sony Interactive Entertainment.",
		},
		{
			Id:          3,
			Name:        "Coca-Cola",
			Quantity:    0,
			Description: "Carbonated soft drink produced by The Coca-Cola Company.",
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v", expected, got)
	}
}

func TestAddDeleteProduct(t *testing.T) {
	initialProducts := getProducts()

	// Add a new product
	var jsonStr = []byte(`{"name":"test","quantity":1,"description":"test"}`)
	_, err := http.Post("http://localhost:8081/stock/add", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("The HTTP request failed with error %s\n", err)
	}

	// Delete the product
	jsonStr = []byte(`{"name":"test"}`)
	_, err = http.Post("http://localhost:8081/stock/delete", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("The HTTP request failed with error %s\n", err)
	}

	var got []Product
	got = getProducts()

	// Check if the product was added and then deleted
	if !reflect.DeepEqual(got, initialProducts) {
		t.Errorf("Expected %+v, got %+v", initialProducts, got)
	}
}

func getProducts() []Product {
	resp, err := http.Get("http://localhost:8081/stock/show")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil
	}

	return products
}
