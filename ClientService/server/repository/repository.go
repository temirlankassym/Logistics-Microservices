package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

type Order struct {
	Id          int32
	ProductName string
	Quantity    int32
	Created     string
	Arrival     string
}

type Repository interface {
	GetOrders(ctx context.Context) ([]Order, error)
	MakeOrder(ctx context.Context, productName string, quantity int32, arrivalDate string) error
}

type Database struct {
	conn *pgx.Conn
}

func Connect(ctx context.Context) (Database, error) {
	config := getConfig()

	conn, err := pgx.Connect(ctx, config)

	if err != nil {
		return Database{}, fmt.Errorf("cannot connect to database")
	}

	return Database{conn: conn}, nil
}

func getConfig() string {
	return fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable",
		"postgres", "mysecretpassword", "5432", "client")
}

func (db *Database) GetOrders(ctx context.Context) ([]Order, error) {
	orders := []Order{}

	rows, err := db.conn.Query(ctx, "Select * from orders")
	if err != nil {
		return []Order{}, err
	}
	fmt.Println("Good")
	for rows.Next() {
		order := Order{}
		err := rows.Scan(&order.Id, &order.ProductName, &order.Quantity, &order.Created, &order.Arrival)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	fmt.Println("Bad")

	return orders, nil
}

func (db *Database) MakeOrder(ctx context.Context, productName string, quantity int32, arrivalDate string) error {
	createdAt := time.Now().Format("2006-01-02 15:04:05")

	_, err := db.conn.Exec(ctx, "INSERT INTO orders (product_name, quantity, created, arrival) VALUES ($1, $2, $3, $4)", productName, quantity, createdAt, arrivalDate)
	if err != nil {
		fmt.Println("Fucking error")
		return err
	}
	return nil
}
