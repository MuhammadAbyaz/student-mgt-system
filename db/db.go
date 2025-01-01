package db

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB{
	err:=godotenv.Load()
	if err != nil{
		fmt.Println("Unable to Connect to DB")
	}
	url := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(url))
	if err != nil {
		fmt.Println("Unable to Connect to DB")
	}
	fmt.Println("Connected To Database")
	return db
}
func Close(dbInstance *gorm.DB){
	sqlDB, err:= dbInstance.DB()
	if err != nil {
		fmt.Println("Failed Closing Database")
	}
	sqlDB.Close()
}