package models

type Person struct {
	Id int `'json:"id"'`
	Name string `'json:"name"'`
	Image string `'json:"image"'`
	BornDate string `'json:"born_date"'`
	BornPlace string `'json:"born_place"'`
	FilmsId []int `'json:"films_id"'`
}

type FilmPerson struct {
	Id int `'json:"id"'`
	Name string `'json:"name"'`
}

type FilmPersons []FilmPerson
