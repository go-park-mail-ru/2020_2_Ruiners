package Models

type Login struct {
	Login    string
	Password string
}

type UserWithoutPassword struct{
	Login string
	Email string
}

type LoginChenge struct {
	Login string
}

type PassChenge struct {
	PasswordOld string
	Password    string
}


type User struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Users = map[string]User{
	"Admin":          User{Login: "Admin", Email: "chekryzhov2000@mail.ru", Password: "Qwerty123456"},
	"AdmiralArkadiy": User{Login: "AdmiralArkadiy", Email: "chekryzhov2000@mail.ru", Password: "Arkadiy1"},
	"ErikDoter":      User{Login: "ErikDoter", Email: "ErikDoter@mail.ru", Password: "commonbaby537"},
}

var Ids = map[string]string{}