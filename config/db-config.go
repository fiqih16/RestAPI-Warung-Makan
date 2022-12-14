package config

import (
	"fmt"
	"os"
	"rumah-makan/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetDBConn() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Isi model
	db.AutoMigrate(&model.Customer{}, &model.Menu{}, &model.Transaction{})
	return db
}

func CloseDBConn(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close database")
	}
	dbSQL.Close()
}