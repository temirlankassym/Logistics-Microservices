package main

import (
	"client/grpcServer/repository"
	"context"
	"fmt"
)

func main() {
	db, err := repository.Connect(context.Background())
	if err != nil {
		fmt.Errorf("can't connect to database")
	}

	orders, _ := db.GetOrders(context.Background())
	fmt.Println(orders)
}
