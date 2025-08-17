package main

import (
	"fmt"

	"gorm.io/gorm/clause"

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

	// Exemplo de lock pessimista com db.Begin()
	tx := db.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	var lockedProduct Product
	// SELECT ... FOR UPDATE
	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&lockedProduct, 1).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("Erro ao buscar produto com lock pessimista:", err)
		return
	}
	lockedProduct.Price += 10
	if err := tx.Save(&lockedProduct).Error; err != nil {
		tx.Rollback()
		fmt.Println("Erro ao salvar produto:", err)
		return
	}
	tx.Commit()
	fmt.Println("Produto atualizado com lock pessimista:", lockedProduct.Name, lockedProduct.Price)

}
