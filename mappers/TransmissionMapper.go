package mappers

import (
	"main/model"
	"main/model/views"
)

func MapToTransmissionViews(Transmissions []*model.Transmission) []views.TransmissionView {
	var result []views.TransmissionView
	for _, x := range Transmissions {
		result = append(result, MapToTransmissionView(x))
	}
	return result
}
func MapToTransmissionView(Transmission *model.Transmission) views.TransmissionView {
	return views.TransmissionView{
		Id:   Transmission.ID,
		Name: Transmission.Name}
}
