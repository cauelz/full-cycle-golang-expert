package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID int `gorm:"primaryKey"`
	Name string
	Price float64
	CategoryID int
	Category Category
	gorm.Model
}

func main() {

	dsn := "user:user_password@tcp(localhost:3306)/my_database?parseTime=True&charset=utf8mb4&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// var category Category
	// category = Category{Name: "Eletrônicos"}
	// db.Create(&category)

	// db.Create(&Product{
	// 	Name: "Dell Inspiron i5",
	// 	Price: 6000.00,
	// 	CategoryID: category.ID,
	// })

	// var products []Product
	// db.Preload("Category").Find(&products)

	// for _, product := range products {
	// 	productJSON, _ := json.Marshal(product)
	// 	fmt.Println("Product: ", string(productJSON))
	// 	fmt.Println("Produto: ", product.Name)
	// 	fmt.Println("Categoria: ", product.Category.Name)
	// }
}