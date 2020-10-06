package main

import (
	"encoding/json"
	"fmt"
	"github.com/lithammer/shortuuid"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Login struct {
	Login    string
	Password string
}

type User struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = map[string]User{
	"Admin":          User{Login: "Admin", Email: "chekryzhov2000@mail.ru", Password: "Qwerty123456"},
	"AdmiralArkadiy": User{Login: "AdmiralArkadiy", Email: "chekryzhov2000@mail.ru", Password: "Arkadiy1"},
	"ErikDoter":      User{Login: "ErikDoter", Email: "ErikDoter@mail.ru", Password: "commonbaby537"},
}

var ids = map[string]string{}

func CreateSession(w http.ResponseWriter, login string) {
	id := shortuuid.New()
	cookie := http.Cookie{
		Name:    "authToken",
		Value:   id,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, &cookie)
	ids[id] = login
}

func main() {
	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/me", isMe)
	http.HandleFunc("/whois", Whois)
	http.HandleFunc("/chengelogin", chengelogin)
	http.HandleFunc("/chengepass", chengepass)
	http.HandleFunc("/chengeavatar", chengeavatar)
	http.HandleFunc("/logout", logoutPage)
	http.HandleFunc("/", mainPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../public"))))
	fmt.Println("starting server at :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("../public/*.html"))
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	l := Login{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if l.Password == users[l.Login].Password && l.Password != "" && l.Login != "" {
		CreateSession(w, l.Login)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func signupPage(w http.ResponseWriter, r *http.Request) {
	u := User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if users[u.Login].Login != u.Login {
		users[u.Login] = u
		//fmt.Println(users)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func isMe(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	login := ids[id.Value]
	if login == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		user := users[login]
		u := &user
		result, _ := json.Marshal(u)
		w.Write(result)
		w.WriteHeader(http.StatusOK)
	}
}

func Whois(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		var login string
		if len(ids) == 0 {
			login = ""
		} else {
			login = ids[id.Value]
			fmt.Println("login who is = ", login)
		}
		fmt.Println(login)
		if login == "" {
			var u User
			u = User{"null", "null", "null"}
			result, _ := json.Marshal(&u)
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		} else {
			user := users[login]
			u := &user
			result, _ := json.Marshal(u)
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}

type LoginChenge struct {
	Login string
}

func chengelogin(w http.ResponseWriter, r *http.Request) {
	l := LoginChenge{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	expiration := time.Now().Add(10 * time.Hour)
	var login string
	if len(ids) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		login = ids[id.Value]
		fmt.Println("login who is = ", login)
	}
	fmt.Println(login)
	user := users[login]

	if user.Login == l.Login {
		fmt.Println("No chenge")
		// http.Redirect(w, r, "/", http.StatusOK)
	} else {
		delete(users, user.Login)
		user.Login = l.Login
		users[l.Login] = user
		fmt.Println(users[l.Login].Login)
		var id = shortuuid.New()
		ids[id] = l.Login
		cookie := http.Cookie{
			Name:     "session_id",
			Value:    id,
			Expires:  expiration,
			HttpOnly: true,
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusOK)
		// http.Redirect(w, r, "/signup", http.StatusBadRequest)
	}
}

type passChenge struct {
	PasswordOld string
	Password    string
}

func chengepass(w http.ResponseWriter, r *http.Request) {
	l := passChenge{}
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(l)
	id, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var login string
	if len(ids) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		login = ids[id.Value]
		fmt.Println("2login who is = ", login)
	}
	fmt.Println(login)
	fmt.Println(users[login].Password + " _ " + l.PasswordOld + " _ " + l.Password)
	if users[login].Password == l.PasswordOld {
		user := users[login]

		delete(users, user.Login)
		user.Password = l.PasswordOld
		users[login] = user
		fmt.Println(users[login].Password)

		http.Redirect(w, r, "/", http.StatusOK)
		// http.Redirect(w, r, "/signup", http.StatusBadRequest)
	}
}

func chengeavatar(w http.ResponseWriter, r *http.Request) {

}

func logoutPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", 401)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	http.Redirect(w, r, "/", http.StatusOK)
}