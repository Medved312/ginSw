package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/database"
	"os"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main/handlers"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.

// @host localhost:5050

func main() {

	logFile, err := os.OpenFile("log.err", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	database.Init()

	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	defer logFile.Close()

	router := gin.Default()

	router.GET("/car/:id", handlers.GetCarHandler)
	router.GET("/mark/:id", handlers.GetMarkHandler)
	router.GET("/model/:id", handlers.GetModelHandler)
	router.GET("/all-mark/", handlers.GetAllMarkHandler)

	router.POST("/car/add/", handlers.CreateCarHandler)
	router.POST("/mark/add/", handlers.CreateMarkHandler)
	router.POST("/model/add/", handlers.CreateModelHandler)

	router.DELETE("/car/:id", handlers.DeleteCarHandler)
	router.DELETE("/mark/:id", handlers.DeleteMarkHandler)
	router.DELETE("/model/:id", handlers.DeleteModelHandler)

	router.PUT("/car/:id", handlers.UpdateCarHandler)
	router.PUT("/mark/:id", handlers.UpdateMarkHandler)
	router.PUT("/model/:id", handlers.UpdateModelHandler)

	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":5050")
}
