package mappers

import (
	"main/model"
	"main/model/views"
)

func MapToProductViews(Products []*model.Product) []views.ProductView {
	var result []views.ProductView

	for _, movie := range Products {
		result = append(result, MapToProductView(movie))
	}
	return result
}
func MapToProductView(Product *model.Product) views.ProductView {
	return views.ProductView{
		Id:          Product.Id,
		Name:        Product.Name,
		Description: Product.Description,
		Categories:  MapToCategoryViews(Product.Categories),
	}
}
