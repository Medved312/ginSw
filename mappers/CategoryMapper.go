package mappers

import (
	"main/model"
	"main/model/views"
)

func MapToCategoryViews(Categories []*model.Category) []views.CategoryView {
	var result []views.CategoryView
	for _, x := range Categories {
		result = append(result, MapToCategoryView(x))
	}
	return result
}
func MapToCategoryView(Category *model.Category) views.CategoryView {
	return views.CategoryView{Id: Category.Id, Name: Category.Name}
}
