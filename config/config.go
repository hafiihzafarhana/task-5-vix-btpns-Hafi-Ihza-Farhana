package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func GetDBConfig() gorm.Dialector {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConfig := databaseConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbName:   os.Getenv("DB_NAME"),
	}

	fmt.Println(dbConfig.dbName, "halo")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		dbConfig.host,
		dbConfig.port,
		dbConfig.user,
		dbConfig.password,
		dbConfig.dbName,
	)

	return postgres.Open(dsn)
}
