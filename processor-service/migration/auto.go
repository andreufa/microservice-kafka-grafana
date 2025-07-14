package main

import (
	"os"
	"processor-service/internal/processor"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	conf := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(conf))
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(&processor.ProcessResult{})
	db.AutoMigrate(&processor.ProcessResult{})
}
