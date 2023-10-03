package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "gorm.io/gorm"
	"log"
	"main/database"
	_ "main/docs"
	"main/mappers"
	"main/model"
	"main/model/saveDTO"
	"net/http"
	"strconv"
)

// @Summary      Get movie
// @Description  Получение фильма
// @Param        id path int true	"id movie"
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.MovieView
// @Router       /movie/{id} [get]
func GetMovieHandler(c *gin.Context) {

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := GetMovieById(uint(id))

	if result == nil {
		log.Println(fmt.Sprintf("фильм с id = %v не удалось найти", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToMovieView(result),
	})
}

func GetMovieById(id uint) *model.Movie {

	var movie *model.Movie

	result := database.GetDB().Model(&model.Movie{}).Preload("Genres").
		Where("id = ?", id).Find(&movie)

	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))
		return nil
	}

	return movie
}

// @Summary      Get genre
// @Description  Получение жанра
// @Param        id path int true	"id genre"
// @Tags         Genres
// @Accept       json
// @Produce      json
// @Success      200  {object}   views.GenreView
// @Router       /genre/{id} [get]
func GetGenreHandler(c *gin.Context) {
	var genre model.Genre

	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := GetGenreById(uint(id))
	if result == nil {
		log.Println(result, fmt.Sprintf("id = %v", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToGenreView(&genre),
	})
}

func GetGenreById(id uint) *model.Genre {
	var genre *model.Genre
	result := database.GetDB().First(&genre, id)
	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))

		return nil
	}

	return genre

}

// @Summary      Get genre
// @Description  Получение списка жанров
// @Tags         Genres
// @Accept       json
// @Produce      json
// @Success      200  {array}   views.GenreView
// @Router       /all-genres [get]
func GetAllGenreHandler(c *gin.Context) {
	var genre []*model.Genre

	result := database.GetDB().Find(&genre)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": mappers.MapToGenreViews(genre),
	})
}

// @Summary      CreateMovie
// @Description  Добавление фильма в базу
// @Param        input body saveDTO.MovieDTO  true  "Title and description of the film"
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Router       /movie/add/ [post]
func CreateMovieHandler(c *gin.Context) {
	var movie *model.Movie
	var createdMovie *saveDTO.MovieDTO
	if err := c.ShouldBindJSON(&createdMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var linkedGenres []*model.Genre
	for _, genreID := range createdMovie.Genres {
		linkedGenres = append(linkedGenres, &model.Genre{Id: genreID})
	}
	movie = &model.Movie{
		Name:        createdMovie.Name,
		Description: createdMovie.Description,
		Genres:      linkedGenres}

	result := database.GetDB().Omit("genres").Create(&movie)
	if result.Error != nil {
		log.Println(result.Error, &movie)
		c.JSON(400, gin.H{
			"message": "Error adding to the database",
		})
		return
	} else {
		newMovie := GetMovieById(movie.Id)
		c.JSON(200, gin.H{
			"result": mappers.MapToMovieView(newMovie),
		})
	}

}

// @Summary      CreateGenre
// @Description  Добавление жанра в базу
// @Param        input body saveDTO.GenreDTO  true  "Сreating a genre"
// @Tags         Genres
// @Accept       json
// @Produce      json
// @Router       /genre/add/ [post]
func CreateGenreHandler(c *gin.Context) {
	var genre *model.Genre

	if err := c.ShouldBindJSON(&genre); err != nil {
		log.Println(err.Error(), &genre)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	row := database.GetDB().Where("name = ?", genre.Name).First(&genre)
	if row.Error != nil {
		result := database.GetDB().Create(&genre)

		if result.Error != nil {
			log.Println(result.Error)
			c.JSON(400, gin.H{
				"message": "Error when adding",
			})
		} else {
			newGenre := GetGenreById(genre.Id)
			c.JSON(200, gin.H{
				"result": mappers.MapToGenreView(newGenre),
			})
		}
	} else {
		c.JSON(409, gin.H{
			"message": "data duplication",
		})
	}
}

// @Summary      DeleteMovie
// @Description  Удаление фильма из базы
// @Param        id path int true	"id movie"
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Router       /movie/{id}  [delete]
func DeleteMovieHandler(c *gin.Context) {
	var movie *model.Movie

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := database.GetDB().Model(&model.Movie{}).Preload("Genres").
		Where("id = ?", id).Find(&movie)
	database.GetDB().Model(movie).Association("Genres").Clear()
	database.GetDB().Delete(&movie)

	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": "Status OK",
	})
}

// @Summary      DeleteGenres
// @Description  Удаление жанра из базы
// @Param        id path int true	"id genre"
// @Tags         Genres
// @Accept       json
// @Produce      json
// @Router       /genre/{id}  [delete]
func DeleteGenreHandler(c *gin.Context) {
	var genre []*model.Genre

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	result := database.GetDB().Model(&model.Genre{}).Preload("Movies").
		Where("id = ?", id).Find(&genre)
	database.GetDB().Model(genre).Association("Movies").Clear()
	database.GetDB().Delete(&genre)

	if result.Error != nil {
		log.Println(result.Error, fmt.Sprintf("id = %v", id))
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": "status OK",
	})
}

// @Summary      UpdateMovie
// @Description  Обновление данных фильма
// @Param		 id path int true	"id movie"
// @Param        input body model.Movie  true  "New values"
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Router       /movie/{id}  [put]
func UpdateMovieHandler(c *gin.Context) {
	var movie *model.Movie = &model.Movie{}
	var upMovie *model.Movie = &model.Movie{}
	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&upMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Where("id = ?", value).First(&movie)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"result": result,
		})
	} else {
		if upMovie.Name != "" {
			movie.Name = upMovie.Name
		}
		if upMovie.Description != "" {
			movie.Description = upMovie.Description
		}
		database.GetDB().Save(&movie)
		c.JSON(200, gin.H{
			"resulst": result,
		})
	}

}

// @Summary      UpdateGenres
// @Description  Обновление данных жанра
// @Param		 id path int true	"id genre"
// @Param        input body model.Genre  true  "New values"
// @Tags         Genres
// @Accept       json
// @Produce      json
// @Router       /genre/{id}  [put]
func UpdateGenreHandler(c *gin.Context) {
	var genre *model.Genre = &model.Genre{}
	var upGenre *model.Genre = &model.Genre{}
	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&upGenre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Where("id = ?", value).First(&genre)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	} else {
		if upGenre.Name != "" {
			genre.Name = upGenre.Name
		}

		database.GetDB().Save(&genre)
		c.JSON(200, gin.H{
			"result": result,
		})
	}
}
