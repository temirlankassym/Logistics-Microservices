package main

import (
	_ "client/docs"
	"client/grpcServer/repository"
	"context"
	"fmt"
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
	orders, _ := db.GetOrders(context.Background())
	fmt.Println(orders)
}
