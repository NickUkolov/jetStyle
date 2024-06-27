package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"notes-service/config"
	_ "notes-service/docs"
	"notes-service/routes"
)

// @title           Notes Service
// @version         1.0
// @description     Notes service API
// @termsOfService  http://swagger.io/terms/

// @host      127.0.0.1:8001
// @BasePath  /

func main() {
	config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()
	routes.RegisterRoutes(r, db)

	port := config.AppConfig.ServerPort
	r.Run(fmt.Sprintf(":%d", port))
}
