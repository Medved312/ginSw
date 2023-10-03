package mappers

import (
	"main/model"
	"main/model/views"
)

func MapToMovieViews(movies []*model.Movie) []views.MovieView {
	var result []views.MovieView

	for _, movie := range movies {
		result = append(result, MapToMovieView(movie))
	}
	return result
}
func MapToMovieView(movie *model.Movie) views.MovieView {
	return views.MovieView{
		Id:          movie.Id,
		Name:        movie.Name,
		Description: movie.Description,
		Genres:      MapToGenreViews(movie.Genres),
	}
}
