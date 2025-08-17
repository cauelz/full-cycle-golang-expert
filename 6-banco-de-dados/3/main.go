package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID int `gorm:"primaryKey"`
	Name string
	Products []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
	ID int `gorm:"primaryKey"`
	Name string
	Price float64
	CategoryID int
	Category Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID int `gorm:"primaryKey"`
	Number string
	ProductID int
}

func main() {

	dsn := "user:user_password@tcp(localhost:3306)/my_database?parseTime=True&charset=utf8mb4&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	var category Category
	category = Category{Name: "Eletronicos"}
	db.Create(&category)

	db.Create(&Product{
		Name: "Dell Inspiron i5",
		Price: 6000.00,
		CategoryID: category.ID,
	})

	db.Create(&Product{
		Name: "Dell Latitude i7",
		Price: 8000.00,
		CategoryID: category.ID,
	})

	db.Create(&SerialNumber{
		Number: "1234",
		ProductID: 1,
	})

	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)

	// for _, product := range products {
	// 	productJSON, _ := json.Marshal(product)
	// 	fmt.Println("Product: ", string(productJSON))
	// 	fmt.Println("Produto: ", product.Name)
	// 	fmt.Println("Categoria: ", product.Category.Name)
	// 	fmt.Println("SerialNumber: ", product.SerialNumber.Number)
	// }

	var categories []Category;

	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		for _, product := range category.Products {
			fmt.Println("Categoria 1: ", category.Name)
			fmt.Println("Product: ", product.Name)
			fmt.Println("SerialNumber: ", product.SerialNumber.Number)
		}
	}
}