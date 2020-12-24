package models

type Film struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Rating      float64 `json:"rating"`
	SumVotes    int     `json:"sum_votes"`
	Description string  `json:"description"`
	MainGenre   string  `json:"main_genre"`
	YoutubeLink string  `json:"youtube_link"`
	BigImg      string  `json:"big_img"`
	SmallImg    string  `json:"small_img"`
	Year        int     `json:"year"`
	Country     string  `json:"country"`
}

type FilmCard struct {
	Id        int     `json:"id"`
	Title     string  `json:"title"`
	MainGenre string  `json:"main_genre"`
	SmallImg  string  `json:"small_img"`
	Year      int     `json:"year"`
	Rating    float64 `json:"rating"`
}

//easyjson:json
type FilmCards []FilmCard

var TranslateGenre = map[string]string{
	"fantasy": "Фантастика",
	"comedy":  "Комедия",
	"horror":  "Ужасы",
	"drama":   "Драма",
	"triller": "Триллер",
	"war":     "Боевик",
}
