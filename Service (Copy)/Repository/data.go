package repository

import (
	"database/sql"
	"fmt"
	"log"
	"service/model"
)

type GetData interface {
	GetUserById(id int) (*model.User, error)
	CreateUser(user model.User) error
	CreateTable() error
	DropTable() error
}

// table creation
//
//	type Tables interface{
//		CreteTable(fields []int ) (error)
//	}
type getData struct {
	db *sql.DB
}

func NewGetData(dba *sql.DB) GetData {
	return &getData{
		db: dba,
	}
}

// type createTable struct{}

// GetUserById returns a user by ID.
//
// It returns a user object and an error.
// If the user is not found, it returns nil and an error.
func (g *getData) GetUserById(id int) (*model.User, error) {
	var user model.User
	query := `SELECT name, email, password, is_verified, active FROM users WHERE id=$1`
	err := g.db.QueryRow(query, id).Scan(&user.Name, &user.Email, &user.Password, &user.IsVerified, &user.Active)
	if err != nil {
		fmt.Println("Failed to execute query")
		return nil, err
	}
	return &user, nil
}

func (g *getData) CreateUser(user model.User) error {
	query := `INSERT INTO users (name, email, password, is_verified, active) VALUES ($1, $2, $3, $4, $5)`
	_, err := g.db.Exec(query, user.Name, user.Email, user.Password, user.IsVerified, user.Active)
	if err != nil {
		fmt.Printf("Error executing insert query: %v\n", err)
		return err
	}
	fmt.Println("User inserted successfully.")
	return nil
}

func (g *getData) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		is_verified BOOLEAN DEFAULT false,
		active BOOLEAN DEFAULT true
	);`
	_, err := g.db.Exec(query)
	if err != nil {
		log.Printf("Failed to create table: %v\n", err)
		return err
	}
	fmt.Println("Table 'users' created or already exists.")
	return nil
}

func (g *getData) DropTable() error {
	query := `DROP TABLE IF EXISTS users;`
	_, err := g.db.Exec(query)
	if err != nil {
		log.Printf("Failed to drop table: %v\n", err)
		return err
	}
	fmt.Println("Table 'users' dropped successfully.")
	return nil
}
