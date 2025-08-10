package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/google/uuid"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(), // Generates a unique ID using github.com/google/uuid
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.ID, p.Name, p.Price)

	return err
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)

	return err
}

func selectProduct(db *sql.DB, id string) (*Product, error) {

	stmt, err := db.Prepare("SELECT * FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var p Product

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product

		err := rows.Scan(&p.ID, &p.Name, &p.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {

	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	dsn := "user:user_password@tcp(localhost:3306)/my_database"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Conexão bem-sucedida!")

	product := NewProduct("Notebook", 3500.00)

	if err := insertProduct(db, product); err != nil {
		log.Fatal(err)
	}

	product.Price = 2500.00

	if err := updateProduct(db, product); err != nil {
		log.Fatal(err)
	}

	// p, err := selectProduct(db, product.ID)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Aqui está o produto buscado: ", p)

	products, err := selectAllProducts(db)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", v.Name, v.Price)
	}

	err = deleteProduct(db, product.ID)

	if err != nil {
		log.Fatal(err)
	}
}