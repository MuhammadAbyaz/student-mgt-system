package db

import (
	"fmt"
	"log"
	"os"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB{
	err:=godotenv.Load()
	if err != nil{
		log.Fatal("Error Loading .env file", err)
	}
	url := os.Getenv("DB_URL")
	if url == "" {
		log.Fatal("DB_URL is not set")
	}
	db, err := gorm.Open(postgres.Open(url))
	if err != nil {
		log.Fatal("Unable to Connect to DB")
	}
	fmt.Println("Connected To Database")
	return db
}
func Close(dbInstance *gorm.DB){
	sqlDB, err:= dbInstance.DB()
	if err != nil {
		log.Fatal("Failed Closing Database")
	}
	sqlDB.Close()
}