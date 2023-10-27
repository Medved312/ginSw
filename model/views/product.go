package views

type ProductView struct {
	Id          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Categories  []CategoryView `json:"Categories"`
}
