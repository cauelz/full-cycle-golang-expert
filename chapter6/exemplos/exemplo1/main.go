package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID string
	Name string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID: uuid.New().String(),
		Name: name,
		Price: price,
	}
}

func main() {
	db, error := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

	if error != nil {
		panic(error)
	}

	defer db.Close()

	product := NewProduct("Macbook m1", 15000.00)

	error = insertProduct(db, product)

	if error != nil {
		panic(error)
	}

}

func insertProduct(db *sql.DB, product *Product) error {

	stmt, error := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")

	if error != nil {
		return error
	}

	defer stmt.Close()

	_, error = stmt.Exec(product.ID, product.Name, product.Price)

	if error != nil {
		return error
	}

	return nil
}