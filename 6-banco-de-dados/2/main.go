package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID int `gorm:"primaryKey"`
	Name string
	Price float64
	gorm.Model
}

func main() {

	dsn := "user:user_password@tcp(localhost:3306)/my_database?parseTime=True&charset=utf8mb4&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// Criando 1 Produto
	// db.Create(&Product{
	// 	Name: "Dell Inspiron", Price: 5000.00,
	// })

	//Criando Vários Produtos
	// db.CreateInBatches(&[]Product{
	// 	{Name: "Dell Inspiron", Price: 5000.00,},
	// 	{Name: "Asus Zenbook i3", Price: 3000.00,},
	// 	{Name: "Asus Zenbook i7", Price: 6000.00,},
	// 	{Name: "Dell ultrabook i5", Price: 5500.00,},
	// }, 1)

	// Selecionando 1 Produto
	// var product Product

	// db.First(&product, 3)
	// fmt.Println("Produto encontrado: ", product)

	// db.First(&product, "name = ?", "Asus Zenbook i3")
	// fmt.Println("Encontrado Asus Zenbook i3: ", product)

	// Select All
	// var products []Product
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println("Product: ", product)
	// }

	// Trabalhando com Limite e Paginação
	// db.Limit(2).Offset(3).Find(&products)
	// for _, product := range products {
	// 	fmt.Println("Product: ", product)
	// }

	// WHERE
	// db.Where("Price > ?", 5000).Find(&products)
	// for _, product := range products {
	// 	fmt.Println("Product: ", product)
	// }

	// WHERE + LIKE
	// db.Where("name LIKE ?", "%i5%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// Atualizando um registro
	// var product Product
	// db.First(&product, 1)
	// product.Name = "Dell Latitude i7"
	// db.Save(&product)

	// db.First(&product, 1)
	// fmt.Println("Produto 1: ", product)

	// Deletando produtos
	// var product Product
	// db.First(&product, 5)
	// fmt.Println("Produto 5: ", product)
	// db.Delete(&product)
	
	// var products []Product
	// db.Save(&Product{
	// 	Name: "Mouse Logitech",
	// 	Price: 300.00,
	// })
	// db.First(&product)
	// fmt.Println("Product 1: ", product)

	// db.Where("ID IN ?", []int{2, 3, 4, 5}).Find(&products).Delete(&products)
}