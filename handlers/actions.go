package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"main/database"
	_ "main/docs"
	"main/model"
	"net/http"
	"strconv"
)

// @Summary      GetListMovies
// @Description  Getting a list of movies
// @Param        limit path int true	"Number of films"
// @Tags         Get
// @Accept       json
// @Produce      json
// @Router       /get/Movie/{limit} [get]
func GetListMovieHandler(c *gin.Context) {
	var movie []*model.Movie
	limit, err := strconv.Atoi(c.Params.ByName("limit"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	row := database.GetDB().Limit(limit).Find(&movie)
	c.JSON(200, gin.H{
		"response": gin.H{
			"methold": http.MethodGet,
			"message": row,
		},
	})
}

// @Summary      Get movie
// @Description  Getting the name of the movie and its description
// @Param        id path int true	"id movie"
// @Tags         Get
// @Accept       json
// @Produce      json
// @Router       /get/movie/{id} [get]
func GetMovieHandler(c *gin.Context) {
	var movie []*model.Movie

	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	row := database.GetDB().First(&movie, value)

	c.JSON(200, gin.H{
		"response": gin.H{
			"methold": http.MethodGet,
			"message": row,
		},
	})
}

// @Summary      Get genre
// @Description  Getting genre
// @Param        id path int true	"id genre"
// @Tags         Get
// @Accept       json
// @Produce      json
// @Router       /get/genre/{id} [get]
func GetGenreHandler(c *gin.Context) {
	var genre []*model.Genre

	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	row := database.GetDB().First(&genre, value)

	c.JSON(200, gin.H{
		"response": gin.H{
			"methold": http.MethodGet,
			"message": row,
		},
	})
}

// @Summary      GetMovieByGenre
// @Description  Getting movies with the selected genre
// @Param        id path int true	"id genre"
// @Tags         Get
// @Accept       json
// @Produce      json
// @Router       /get/movieByGenres/{id} [get]
func GetMovieByGenre(c *gin.Context) {
	var movieGenre []*model.Movie_genre

	value, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	row := database.GetDB().Where("id_genre = ?", value).Find(&movieGenre)
	c.JSON(200, gin.H{
		"response": gin.H{
			"methold": http.MethodGet,
			"message": row,
		},
	})
}

// @Summary      CreateMovie
// @Description  Сreating a movie record
// @Param        input body model.Movie  true  "Title and description of the film"
// @Tags         Create
// @Accept       json
// @Produce      json
// @Router       /movie/add/ [post]
func CreateMovieHandler(c *gin.Context) {
	var movie *model.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := movie.Name_movie
	row := database.GetDB().Where("name_movie = ?", name).First(&movie)
	if row.Error != nil {
		result := database.Add(&movie)

		if result != nil {
			c.JSON(400, gin.H{
				"response": gin.H{
					"message": result.Error,
				},
			})
		} else {
			c.JSON(200, gin.H{
				"response": gin.H{
					"methold": http.MethodPost,
					"message": result,
				},
			})
		}
	} else {
		c.JSON(500, gin.H{
			"response": gin.H{
				"methold": http.MethodPost,
				"message": "data duplication",
			},
		})
	}

}

// @Summary      CreateGenre
// @Description  Сreating a movie record
// @Param        input body model.Genre  true  "Сreating a genre"
// @Tags         Create
// @Accept       json
// @Produce      json
// @Router       /genre/add/ [post]
func CreateGenreHandler(c *gin.Context) {
	var genre *model.Genre

	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := genre.Name_genre
	row := database.GetDB().Where("name_genre = ?", name).First(&genre)

	if row.Error != nil {
		result := database.Add(&genre)

		if result.Error != nil {
			c.JSON(404, gin.H{
				"response": gin.H{
					"methold": http.MethodPost,
					"message": result.Error,
				},
			})
		} else {
			c.JSON(200, gin.H{
				"response": gin.H{
					"methold": http.MethodPost,
					"message": result,
				},
			})
		}
	} else {
		c.JSON(500, gin.H{
			"response": gin.H{
				"methold": http.MethodPost,
				"message": "data duplication",
			},
		})
	}
}

// @Summary      CreateMovieGenres
// @Description  Сreating a movie record
// @Param        input body model.Movie_genre  true  "Linking genres to a film"
// @Tags         Create
// @Accept       json
// @Produce      json
// @Router       /movieGenres/add/ [post]
func CreateMovieGenresHandler(c *gin.Context) {
	var movieGenres *model.Movie_genre

	if err := c.ShouldBindJSON(&movieGenres); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := movieGenres.Id_movie

	row := database.GetDB().Where("name_genre = ?", name).First(&movieGenres)

	if row.Error != nil {
		result := database.Add(&movieGenres)

		if result.Error != nil {
			c.JSON(404, gin.H{
				"response": gin.H{
					"message": result.Error,
				},
			})
		} else {
			c.JSON(200, gin.H{
				"response": gin.H{
					"methold": http.MethodPost,
					"message": result,
				},
			})
		}
	} else {
		c.JSON(500, gin.H{
			"response": gin.H{
				"message": "data duplication",
			},
		})
	}
}

// @Summary      DeleteMovie
// @Description  Deleting data from a table
// @Param        id path int true	"id movie"
// @Tags         Delete
// @Accept       json
// @Produce      json
// @Router       /movie/delete/{id}  [post]
func DeleteMovieHandler(c *gin.Context) {
	var movie []*model.Movie

	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	row := database.GetDB().Delete(&movie, value)

	c.JSON(200, gin.H{
		"response": gin.H{
			"methold": http.MethodPost,
			"message": row,
		},
	})
}

// @Summary      DeleteGenres
// @Description  Deleting data from a table
// @Param        id path int true	"id genre"
// @Tags         Delete
// @Accept       json
// @Produce      json
// @Router       /genre/delete/{id}  [post]
func DeleteGenreHandler(c *gin.Context) {
	var genre []*model.Genre

	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	row := database.GetDB().First(&genre, value)

	c.JSON(200, gin.H{
		"response": gin.H{
			"methold": http.MethodPost,
			"message": row,
		},
	})
}

// @Summary      DeleteMovieGenres
// @Description  Deleting data from a table
// @Param        id path int true	"id movieGenres"
// @Tags         Delete
// @Accept       json
// @Produce      json
// @Router       /movieGenres/delete/{id}  [post]
func DeleteMovieGenresHandler(c *gin.Context) {
	var mGenre []*model.Movie_genre

	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	row := database.GetDB().First(&mGenre, value)

	c.JSON(200, gin.H{
		"response": gin.H{
			"methold": http.MethodPost,
			"message": row,
		},
	})
}

// @Summary      UpdateMovie
// @Description  Updating values in the table movie
// @Param		 id path int true	"id movie"
// @Param        input body model.Movie  true  "New values"
// @Tags         Update
// @Accept       json
// @Produce      json
// @Router       /movie/update/{id}  [post]
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
			"response": gin.H{
				"message": result,
			},
		})
	} else {
		if upMovie.Name_movie != "" {
			movie.Name_movie = upMovie.Name_movie
		}
		if upMovie.Description != "" {
			movie.Description = upMovie.Description
		}
		database.GetDB().Save(&movie)
		c.JSON(200, gin.H{
			"response": gin.H{
				"methold": http.MethodPost,
				"message": result,
			},
		})
	}

}

// @Summary      UpdateGenres
// @Description  Updating values in the table genre
// @Param		 id path int true	"id genre"
// @Param        input body model.Genre  true  "New values"
// @Tags         Update
// @Accept       json
// @Produce      json
// @Router       /genre/update/{id}  [post]
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
		c.JSON(404, gin.H{
			"response": gin.H{
				"message": result,
			},
		})
	} else {
		if upGenre.Name_genre != "" {
			genre.Name_genre = upGenre.Name_genre
		}

		database.GetDB().Save(&genre)
		c.JSON(200, gin.H{
			"response": gin.H{
				"methold": http.MethodPost,
				"message": result,
			},
		})
	}

}

// @Summary      UpdateMovieGenres
// @Description  Updating values in the table movieGenres
// @Param		 id path int true	"id movie_genre"
// @Param        input body model.Movie_genre  true  "New values"
// @Tags         Update
// @Accept       json
// @Produce      json
// @Router       /movieGenres/update/{id} [post]
func UpdateMovieGenresHandler(c *gin.Context) {
	var movieGenre *model.Movie_genre = &model.Movie_genre{}
	var upMovieGenre *model.Movie_genre = &model.Movie_genre{}
	value, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&upMovieGenre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.GetDB().Where("id = ?", value).First(&movieGenre)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"response": gin.H{
				"message": result,
			},
		})
	} else {
		if upMovieGenre.Id_genre != "" {
			movieGenre.Id_genre = upMovieGenre.Id_genre
		}
		if upMovieGenre.Id_movie != "" {
			movieGenre.Id_movie = upMovieGenre.Id_movie
		}
		database.GetDB().Save(&movieGenre)
		c.JSON(200, gin.H{
			"response": gin.H{
				"methold": http.MethodPost,
				"message": result,
			},
		})
	}

}
