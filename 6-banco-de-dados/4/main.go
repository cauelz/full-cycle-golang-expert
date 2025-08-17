package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID int `gorm:"primaryKey"`
	Name string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID int `gorm:"primaryKey"`
	Name string
	Price float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}


func main() {
	dsn := "user:user_password@tcp(localhost:3306)/my_database?parseTime=True&charset=utf8mb4&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	var category Category
	category = Category{Name: "Cozinha"}
	db.Create(&category)

	var category2 Category
	category2 = Category{Name: "Eletronicos"}
	db.Create(&category2)

	db.Create(&Product{
		Name: "Micro Ondas",
		Price: 600.00,
		Categories: []Category{category, category2},
	})

	db.Create(&Product{
		Name: "Geladeira",
		Price: 600.00,
		Categories: []Category{category, category2},	
	})

	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	for _, category := range categories {
		fmt.Println("Category: ", category.Name)
		for _, product := range category.Products {
			fmt.Println("Product: ", product.Name)
		}
	}
}
