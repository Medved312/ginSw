package model

type Movie struct {
	Id          uint   `json:"id"`
	Name_movie  string `json:"name_movie"`
	Description string `json:"description"`
}
type Genre struct {
	Id         uint   `json:"id"`
	Name_genre string `json:"name_genre"`
}
type Movie_genre struct {
	Id       uint   `json:"id"`
	Id_movie string `json:"id_movie"`
	Id_genre string `json:"id_genre"`
}
