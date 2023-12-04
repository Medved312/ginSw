package mappers

import (
	"main/model"
	"main/model/views"
)

func MapToModelViews(Models []*model.Model) []views.ModelView {
	var result []views.ModelView
	for _, x := range Models {
		result = append(result, MapToModelView(x))
	}
	return result
}
func MapToModelView(Model *model.Model) views.ModelView {
	return views.ModelView{
		Id:   Model.ID,
		Name: Model.Name}
}
