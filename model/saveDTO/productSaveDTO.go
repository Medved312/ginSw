package saveDTO

type ProductDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Categories  []uint `json:"idCategories"`
}
