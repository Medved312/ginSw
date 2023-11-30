package mappers

import (
	"main/model"
	"main/model/views"
)

func MapToCarViews(Cars []*model.Car) []views.CarView {
	var result []views.CarView

	for _, car := range Cars {
		result = append(result, MapToCarView(car))
	}
	return result
}
func MapToCarView(Car *model.Car) views.CarView {
	return views.CarView{
		Id:             Car.ID,
		Price:          Car.Price,
		Power:          Car.Power,
		Consumption:    Car.Consumption,
		Overclocking:   Car.Overclocking,
		MaxSpeed:       Car.MaxSpeed,
		ModelID:        Car.ModelID,
		TransmissionID: Car.TransmissionID,
		Description:    Car.Description,
	}
}
