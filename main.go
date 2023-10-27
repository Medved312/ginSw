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

	router.GET("/product/:id", handlers.GetProductHandler)
	router.GET("/category/:id", handlers.GetCategoryHandler)
	router.GET("/all-categories/", handlers.GetAllCategoryHandler)

	router.POST("/product/add/", handlers.CreateProductHandler)
	router.POST("/category/add/", handlers.CreateCategoryHandler)

	router.DELETE("/product/:id", handlers.DeleteProductHandler)
	router.DELETE("/category/:id", handlers.DeleteGenreHandler)

	router.PUT("/product/:id", handlers.UpdateProductHandler)
	router.PUT("/category/:id", handlers.UpdateCategoryHandler)

	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":5050")
}
