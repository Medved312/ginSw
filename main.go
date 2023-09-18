package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main/database"
	"main/handlers"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.

// @host localhost:5050

func main() {
	database.Init()
	router := gin.Default()

	router.GET("/get/Movie/:limit", handlers.GetListMovieHandler)
	router.GET("/get/movie/:id", handlers.GetMovieHandler)
	router.GET("/get/genre/:id", handlers.GetGenreHandler)
	router.GET("/get/movieByGenres/:id", handlers.GetMovieByGenre)

	router.POST("/movie/add/", handlers.CreateMovieHandler)
	router.POST("/genre/add/", handlers.CreateGenreHandler)
	router.POST("/movieGenres/add/", handlers.CreateMovieGenresHandler)

	router.POST("/movie/delete/:id", handlers.DeleteMovieHandler)
	router.POST("/genre/delete/:id", handlers.DeleteGenreHandler)
	router.POST("/movieGenres/delete/:id", handlers.DeleteMovieGenresHandler)

	router.POST("/movie/update/:id", handlers.UpdateMovieHandler)
	router.POST("/genre/update/:id", handlers.UpdateGenreHandler)
	router.POST("/movieGenres/update/:id", handlers.UpdateMovieGenresHandler)

	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":5050")
}
