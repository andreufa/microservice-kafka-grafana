package main

import (
	"filter-service/internal/raw"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("./filter-service/.env")
	if err != nil {
		panic(err)
	}
	conf := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(conf))
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(&raw.RawData{})
	db.AutoMigrate(&raw.RawData{})
}
