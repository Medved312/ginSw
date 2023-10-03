package main

import (
	"github.com/gin-gonic/gin"
	"log"
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

	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	defer logFile.Close()

	router := gin.Default()

	router.GET("/movie/:id", handlers.GetMovieHandler)
	router.GET("/genre/:id", handlers.GetGenreHandler)
	router.GET("/all-genres/", handlers.GetAllGenreHandler)

	router.POST("/movie/add/", handlers.CreateMovieHandler)
	router.POST("/genre/add/", handlers.CreateGenreHandler)

	router.DELETE("/movie/:id", handlers.DeleteMovieHandler)
	router.DELETE("/genre/:id", handlers.DeleteGenreHandler)

	router.PUT("/movie/:id", handlers.UpdateMovieHandler)
	router.PUT("/genre/:id", handlers.UpdateGenreHandler)

	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":5050")
}
