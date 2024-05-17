package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type User struct {
	Id    int32
	Name  string
	Email string
}

func Connect(ctx context.Context) (*pgx.Conn, error) {
	config := getConfig()

	db, err := pgx.Connect(ctx, config)

	if err != nil {
		return nil, fmt.Errorf("cannot connect to database")
	}

	return db, nil
}

func getConfig() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return fmt.Sprintf("postgresql://%s:%s@localhost:%s?sslmode=disable",
		os.Getenv("DBNAME"), os.Getenv("PASSWORD"), os.Getenv("PORT"))
}

func GetUser(ctx context.Context, db *pgx.Conn, id int32) (User, error) {
	user := User{}

	err := db.QueryRow(ctx, "Select * from users where id = $1", id).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func AddUser(ctx context.Context, db *pgx.Conn, name string, email string) (int32, error) {
	var id int32

	err := db.QueryRow(ctx, "Insert into users (name, email) values ($1, $2) returning id", name, email).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ListUsers(ctx context.Context, db *pgx.Conn) ([]User, error) {
	users := []User{}

	rows, err := db.Query(ctx, "Select * from users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
