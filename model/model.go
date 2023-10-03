package model

type Movie struct {
	Id          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name_movie"`
	Description string   `json:"description"`
	Genres      []*Genre `gorm:"many2many:movie_genres"`
}
type Genre struct {
	Id     uint     `json:"id" gorm:"primaryKey"`
	Name   string   `json:"name_genre"`
	Movies []*Movie `gorm:"many2many:movie_genres"`
}
