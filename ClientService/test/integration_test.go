package test

import (
	"net/http"
	"testing"
)

func TestShowOrders(t *testing.T) {
	resp, err := http.Get("http://138.197.180.252:8080/orders/show")
	if err != nil {
		t.Fatalf("The HTTP request failed with error %s\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
