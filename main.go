package main

import (
	"go-api-questdb/controllers"
	"go-api-questdb/models"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()

	dbConfig := models.DBConfig{
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name: os.Getenv("DB_NAME"),
		Port: os.Getenv("DB_PORT"),
	}
	err := dbConfig.ConnectDatabase()

	if err != nil {
		panic(err)
	}
	
	r.POST("/trackers", controllers.CreateTracker)
	r.GET("/trackers", controllers.GetTrackers)

	r.Run("localhost:3000")
}
