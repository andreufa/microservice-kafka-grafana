package main

import (
	"filter-service/internal/raw"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	conf := "host=localhost user=postgres password=fservice dbname=fservice_db port=15433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(conf))
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(&raw.RawData{})
	db.AutoMigrate(&raw.RawData{})
}
