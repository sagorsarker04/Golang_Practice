package main

import (
	"database/sql"
	"fmt"
	"gorom/gorom"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func (User) TableName() string {
	return "users"
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func (Product) TableName() string {
	return "products"
}

func main() {
	dsn := "postgres://postgres:123@localhost:5432/TestDB?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}

	// Get all users
	userRepo := gorom.NewRepository[User](db)
	users, err := userRepo.All()
	if err != nil {
		log.Fatal("Fetch users error:", err)
	}

	fmt.Println("Users:")
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	}

	// Get all products
	productRepo := gorom.NewRepository[Product](db)
	products, err := productRepo.All()
	if err != nil {
		log.Fatal("Fetch products error:", err)
	}

	fmt.Println("\nProducts:")
	for _, p := range products {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}
