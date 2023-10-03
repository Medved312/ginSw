package saveDTO

type MovieDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Genres      []uint `json:"id_genres"`
}
