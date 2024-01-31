package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Define a model struct
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB

func main() {
	// Replace 'your_mysql_connection_string' with your actual MySQL connection string
	dsn := "root:Polar@12345(localhost:3306)/new_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	defer func() {
		// Close the GORM database connection
		if err := db.Close(); err != nil {
			fmt.Println("Error closing database:", err)
		}
	}()

	// Auto Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "P1", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                  // find product with id 1
	db.First(&product, "code = ?", "P1")   // find product with code P1

	// Update
	db.Model(&product).Update("Price", 200)

	// Delete
	db.Delete(&product, 1)
}
