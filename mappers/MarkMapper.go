package mappers

import (
	"main/model"
	"main/model/views"
)

func MapToMarkViews(Marks []*model.Mark) []views.MarkView {
	var result []views.MarkView
	for _, x := range Marks {
		result = append(result, MapToMarkView(x))
	}
	return result
}
func MapToMarkView(Mark *model.Mark) views.MarkView {
	return views.MarkView{
		Id:      Mark.ID,
		Name:    Mark.Name,
		ModelID: Mark.ModelID}
}
