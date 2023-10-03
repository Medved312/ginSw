package views

type MovieView struct {
	Id          uint        `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Genres      []GenreView `json:"genres"`
}
