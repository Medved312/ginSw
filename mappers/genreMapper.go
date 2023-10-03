package mappers

import (
	"main/model"
	"main/model/views"
)

func MapToGenreViews(genres []*model.Genre) []views.GenreView {
	var result []views.GenreView
	for _, x := range genres {
		result = append(result, MapToGenreView(x))
	}
	return result
}
func MapToGenreView(genre *model.Genre) views.GenreView {
	return views.GenreView{Id: genre.Id, Name: genre.Name}
}
