package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type Product struct {
	Id          int32
	ProductName string
	Quantity    int32
	Description string
}

type Repository interface {
	GetProduct(ctx context.Context, productName string) (Product, error)
	DecrementProduct(ctx context.Context, productName string, quantity int32, c chan int32) error
	ShowStock(ctx context.Context) ([]Product, error)
	AddProduct(ctx context.Context, productName string, quantity int32, description string) error
	DeleteProduct(ctx context.Context, productName string) error
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
		"postgres", "mysecretpassword", "5432", "inventory")
}

func (db *Database) GetProduct(ctx context.Context, productName string) (Product, error) {
	var product Product

	row := db.conn.QueryRow(ctx, "SELECT * FROM products WHERE product_name = $1", productName)

	err := row.Scan(&product.Id, &product.ProductName, &product.Quantity, &product.Description)

	if err != nil {
		return Product{}, fmt.Errorf("cannot get product")
	}

	return product, nil
}

func (db *Database) DecrementProduct(ctx context.Context, productName string, quantity int32, c chan int32) error {
	var stock int32
	var productCount int32

	transaction, err := db.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error when starting transaction %w", err)
	}

	row := db.conn.QueryRow(ctx, "SELECT quantity FROM products WHERE product_name = $1", productName)
	err = row.Scan(&stock)
	if err != nil {
		return fmt.Errorf("there is no such product")
	}

	// writing to channel number of missing product quantity
	c <- quantity - stock

	// if not enough products in stock give what is present and set quantity to 0
	if stock < quantity {
		productCount = 0
	} else {
		productCount = stock - quantity
	}

	_, err = db.conn.Exec(ctx, "UPDATE products SET quantity = $1 WHERE product_name = $2", productCount, productName)

	if err != nil {
		transaction.Rollback(ctx)
		return fmt.Errorf("cannot decrement product")
	}

	if err = transaction.Commit(ctx); err != nil {
		transaction.Rollback(ctx)
		return fmt.Errorf("can't commit transaction: %w", err)
	}

	return nil
}

func (db *Database) ShowStock(ctx context.Context) ([]Product, error) {
	rows, err := db.conn.Query(ctx, "SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("cannot get products")
	}

	var products []Product

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.Id, &product.ProductName, &product.Quantity, &product.Description)
		if err != nil {
			return nil, fmt.Errorf("cannot scan product")
		}
		products = append(products, product)
	}

	return products, nil
}

func (db *Database) AddProduct(ctx context.Context, productName string, quantity int32, description string) error {
	_, err := db.conn.Exec(ctx, "INSERT INTO products (product_name, quantity, description) VALUES ($1, $2, $3)", productName, quantity, description)
	if err != nil {
		return fmt.Errorf("cannot add product")
	}

	return nil
}

func (db *Database) DeleteProduct(ctx context.Context, productName string) error {
	_, err := db.conn.Exec(ctx, "DELETE FROM products WHERE product_name = $1", productName)
	if err != nil {
		return fmt.Errorf("cannot delete product")
	}

	return nil
}
