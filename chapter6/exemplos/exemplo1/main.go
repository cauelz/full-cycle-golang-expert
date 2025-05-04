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

	product.Name = "Macbook m2"

	error = updateProduct(db, product)

	if error != nil {
		panic(error)
	}

	// product, error = selectProduct(db, product.ID)

	// if error != nil {
	// 	panic(error)
	// }

	// fmt.Printf("O produto %s tem o preço de R$ %.2f", product.Name, product.Price)

	error = deleteProduct(db, product.ID)

	if error != nil {
		panic(error)
	}

	products, error := selectAllProducts(db)

	if error != nil {
		panic(error)
	}

	for _, product := range products {
		println(product.ID, product.Name, product.Price)
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

func updateProduct(db *sql.DB, product *Product) error {
	stmt, error := db.Prepare("update products set name = ?, price = ? where id = ?")

	if error != nil {
		return error
	}

	defer stmt.Close()

	_, error = stmt.Exec(product.Name, product.Price, product.ID)

	if error != nil {
		return error
	}

	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {

	stmt, error := db.Prepare("select id, name, price from products where id = ?")

	if error != nil {
		return nil, error
	}

	defer stmt.Close()

	var product Product

	error = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if error != nil {
		return nil, error
	}

	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]*Product, error) {

	rows, error := db.Query("select id, name, price from products")

	if error != nil {
		return nil, error
	}

	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var product Product

		error = rows.Scan(&product.ID, &product.Name, &product.Price)

		if error != nil {
			return nil, error
		}

		products = append(products, &product)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, error := db.Prepare("delete from products where id = ?")

	if error != nil {
		return error
	}

	defer stmt.Close()

	_, error = stmt.Exec(id)

	if error != nil {
		return error
	}

	return nil
}