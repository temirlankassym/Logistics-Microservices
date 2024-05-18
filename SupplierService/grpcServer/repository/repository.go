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

type Supplier struct {
	Id          int32  `json:"id"`
	ProductName string `json:"product_name"`
	CompanyName string `json:"company_name"`
	DaysToShip  int32  `json:"days_to_ship"`
}

func (db *Database) ShowSuppliers(ctx context.Context) ([]Supplier, error) {
	rows, err := db.conn.Query(ctx, "SELECT * FROM company_products")
	if err != nil {
		return nil, fmt.Errorf("cannot get suppliers")
	}

	var suppliers []Supplier
	for rows.Next() {
		var supplier Supplier
		err = rows.Scan(&supplier.Id, &supplier.ProductName, &supplier.CompanyName, &supplier.DaysToShip)
		if err != nil {
			return nil, fmt.Errorf("cannot get suppliers")
		}
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (db *Database) AddSupplier(ctx context.Context, productName string, companyName string, daysToShip int32) error {
	_, err := db.conn.Exec(ctx, "INSERT INTO company_products (product_name, company_name, days_to_ship) VALUES ($1, $2, $3)", productName, companyName, daysToShip)

	if err != nil {
		return fmt.Errorf("cannot add supplier")
	}

	return nil
}

func (db *Database) DeleteSupplier(ctx context.Context, productName string) error {
	_, err := db.conn.Exec(ctx, "DELETE FROM company_products WHERE product_name = $1", productName)

	if err != nil {
		return fmt.Errorf("cannot delete supplier")
	}

	return nil
}

type Delivery struct {
	Id          int32  `json:"id"`
	ProductName string `json:"product_name"`
	Quantity    int32  `json:"quantity"`
	Arrival     string `json:"arrival"`
}

func (db *Database) ShowDeliveries(ctx context.Context) ([]Delivery, error) {
	rows, err := db.conn.Query(ctx, "SELECT * FROM deliveries")
	if err != nil {
		return nil, fmt.Errorf("cannot get deliveries")
	}

	var deliveries []Delivery
	for rows.Next() {
		var delivery Delivery
		err = rows.Scan(&delivery.Id, &delivery.ProductName, &delivery.Quantity, &delivery.Arrival)
		if err != nil {
			return nil, fmt.Errorf("cannot get deliveries")
		}
		deliveries = append(deliveries, delivery)
	}

	return deliveries, nil
}
