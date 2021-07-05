package main

import (
	"fmt"

	"github.com/adiputra22/golang-gin-first/Config"
	"github.com/adiputra22/golang-gin-first/Models"
	"github.com/adiputra22/golang-gin-first/Routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	// connect db
	connectDB()

	r := Routes.SetupRouter()
	//running
	r.Run(":8000")
}

func connectDB() {
	dbConfig := Config.BuildDBConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port)
	Config.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}
	Config.DB.AutoMigrate(&Models.User{})
}
