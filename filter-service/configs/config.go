package configs

import (
	"filter-service/commom"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(commom.ErrLoadConfig, err)
	}
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}
