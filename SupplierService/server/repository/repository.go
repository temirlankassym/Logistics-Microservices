package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	GetDeliveryTime(ctx context.Context, productName string) (int32, error)
	CreateDelivery(ctx context.Context, productName string, quantity int32, arrival string) error
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
		"postgres", "mysecretpassword", "5432", "supplier")
}

func (db *Database) GetDeliveryTime(ctx context.Context, productName string) (int32, error) {
	var days int32
	err := db.conn.QueryRow(ctx, "SELECT days_to_ship FROM company_products WHERE product_name = $1", productName).Scan(&days)

	if err != nil {
		return 0, fmt.Errorf("cannot get product")
	}

	return days, nil
}

func (db *Database) CreateDelivery(ctx context.Context, productName string, quantity int32, arrival string) error {
	_, err := db.conn.Exec(ctx, "INSERT INTO deliveries (product_name, quantity, arrival) VALUES ($1, $2, $3)", productName, quantity, arrival)

	if err != nil {
		return fmt.Errorf("cannot create delivery")
	}

	return nil
}
