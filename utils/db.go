package utils

import (
	"fmt"
	"log"
	"os"

	"../models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //Gorm sqlite dialect interface

	//	_ "github.com/jinzhu/gorm/dialects/postgres" //Gorm postgres dialect interface
	"github.com/joho/godotenv"
)

//ConnectDB function: Make database connection
func ConnectDB() *gorm.DB {

	//Load environmenatal variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// username := os.Getenv("databaseUser")
	// password := os.Getenv("databasePassword")
	// databaseName := os.Getenv("databaseName")
	// databaseHost := os.Getenv("databaseHost")

	//Define DB connection string
	//dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)

	//connect to db URI
	db, err := gorm.Open("sqlite3", "./"+os.Getenv("sqliteDbName"))

	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	// close db when not in use
	// defer db.Close()

	// Migrate the schema
	db.AutoMigrate(
		&models.User{},
		&models.Meal{},
		&models.EatenMeal{},
		&models.MealRating{},
		&models.Ingredient{},
		&models.Person{})

	fmt.Println("Successfully connected!", db)
	return db
}
