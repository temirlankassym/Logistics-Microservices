package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type Supplier struct {
	Id          int32  `json:"id"`
	ProductName string `json:"product_name"`
	CompanyName string `json:"company_name"`
	DaysToShip  int32  `json:"days_to_ship"`
}

func TestShowSuppliers(t *testing.T) {
	var got []Supplier

	got = getSuppliers()

	expected := []Supplier{
		{
			Id:          1,
			ProductName: "iPhone",
			CompanyName: "Apple Inc.",
			DaysToShip:  3,
		},
		{
			Id:          2,
			ProductName: "PlayStation",
			CompanyName: "Sony Interactive Entertainment",
			DaysToShip:  5,
		},
		{
			Id:          3,
			ProductName: "Coca-Cola",
			CompanyName: "The Coca-Cola Company",
			DaysToShip:  2,
		},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v", expected, got)
	}
}

func TestAddDeleteSupplier(t *testing.T) {
	initialSuppliers := getSuppliers()

	// Add a new supplier
	var jsonStr = []byte(`{"name":"test","company_name":"test company","days_to_ship":1}`)
	_, err := http.Post("http://104.248.21.144:8082/suppliers/add", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("The HTTP request failed with error %s\n", err)
	}

	// Delete the supplier
	jsonStr = []byte(`{"name":"test"}`)
	_, err = http.Post("http://104.248.21.144:8082/suppliers/delete", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("The HTTP request failed with error %s\n", err)
	}

	var got []Supplier
	got = getSuppliers()

	// Check if the supplier was added and then deleted
	if !reflect.DeepEqual(got, initialSuppliers) {
		t.Errorf("Expected %+v, got %+v", initialSuppliers, got)
	}
}

func TestShowDeliveries(t *testing.T) {
	resp, err := http.Get("http://104.248.21.144:8082/deliveries/show")
	if err != nil {
		t.Fatalf("The HTTP request failed with error %s\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func getSuppliers() []Supplier {
	resp, err := http.Get("http://104.248.21.144:8082/suppliers/show")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var suppliers []Supplier
	err = json.Unmarshal(data, &suppliers)
	if err != nil {
		return nil
	}

	return suppliers
}
