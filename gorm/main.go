package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	createPostgresdb()
}

func createSqlite3db() {
	// Open db and get reference to db using orm
	db, error := gorm.Open("sqlite3", "test.db")
	if error != nil {
		fmt.Println(error)
		panic("failed to connect database.")
	}

	// Using defer close db after all operations are complete
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1) // Find product with id 1
	db.First(&product, "code = ?", "L1212")

	// Update - update's product price to 2000
	db.Model(&product).Update("Price", 3000)


	// Delete - delete product
	db.Delete(&product)
}

func createPostgresdb()  {

	// step 1 : open db
	db, error := gorm.Open("postgres", "host=localhost user=gormtest dbname=gormtest sslmode=disable password=gormtestpassword")
	if error != nil {
		fmt.Println(error)
		panic("failed to connect database.")
	}

	defer db.Close()

	// step 2 : migrate the schema
	db.AutoMigrate(&Product{})

	// step 3 : create
	//db.Create(&Product{Code: "P1212", Price: 5000})

	// step 4 : read
	var product Product
	db.First(&product, 1) // Find product with ID 1
	db.First(&product, "code = ?", "P1212")
	fmt.Println(product)

	// step 5 : update
	db.Model(db.First(&product, "code = ?", "P1212")).Update("Price", 65000)
	fmt.Println(product)

	// step 6 : delete
	//db.Delete(&product)
	
}