package Handlers

import (
	"../Models"
	"encoding/json"
	"fmt"
	"github.com/lithammer/shortuuid"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func CreateSession(w http.ResponseWriter, login string) {
	id := shortuuid.New()
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   id,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, &cookie)
	Models.Ids[id] = login
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("../public/*.html"))
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	l := Models.Login{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if l.Password == Models.Users[l.Login].Password && l.Password != "" && l.Login != "" {
		CreateSession(w, l.Login)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	u := Models.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if Models.Users[u.Login].Login != u.Login {
		Models.Users[u.Login] = u
		CreateSession(w, u.Login)
		w.WriteHeader(http.StatusOK)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func IsMe(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("session_id")
	if err != nil {
		user := Models.UserWithoutPassword{Login: "", Email: ""}
		u := &user
		result, err := json.Marshal(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
		return
	}
	login := Models.Ids[id.Value]
	if login == "" {
		user := Models.UserWithoutPassword{Login: "", Email: ""}
		u := &user
		result, err := json.Marshal(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(result)
	} else {
		user := Models.Users[login]
		u := Models.UserWithoutPassword{Login: user.Login, Email: user.Email}
		uPtr := &u
		result, err := json.Marshal(uPtr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(result)
		w.WriteHeader(http.StatusOK)
	}
}

func Whois(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("session_id")
	if err != nil {
		var u Models.UserWithoutPassword
		u = Models.UserWithoutPassword{Login: "null", Email: "null"}
		result, err := json.Marshal(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(result)
		return
	} else {
		var login string
		if len(Models.Ids) == 0 {
			login = ""
		} else {
			login = Models.Ids[id.Value]
		}
		if login == "" {
			var u Models.UserWithoutPassword
			u = Models.UserWithoutPassword{Login: "null", Email: "null"}
			result, err := json.Marshal(&u)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		} else {
			user := Models.Users[login]
			u := Models.UserWithoutPassword{Login: user.Login, Email: user.Email}
			fmt.Println(u)
			uPtr := &u
			result, err := json.Marshal(uPtr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(result)
		}
	}
}



func Chengelogin(w http.ResponseWriter, r *http.Request) {
	l := Models.LoginChenge{}
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
	if len(Models.Ids) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		login = Models.Ids[id.Value]
	}
	user := Models.Users[login]

	if user.Login == l.Login {
		// http.Redirect(w, r, "/", http.StatusOK)
	} else {
		delete(Models.Users, user.Login)
		user.Login = l.Login
		Models.Users[l.Login] = user
		var id = shortuuid.New()
		Models.Ids[id] = l.Login
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

func Chengepass(w http.ResponseWriter, r *http.Request) {
	l := Models.PassChenge{}
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
	var login string
	if len(Models.Ids) == 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		login = Models.Ids[id.Value]
	}
	if Models.Users[login].Password == l.PasswordOld {
		user := Models.Users[login]

		delete(Models.Users, user.Login)
		user.Password = l.Password
		Models.Users[login] = user
		http.Redirect(w, r, "/", http.StatusOK)
		// http.Redirect(w, r, "/signup", http.StatusBadRequest)
	}
}

func Changeavatar(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	Models.ImgId++
	s := strconv.Itoa(Models.ImgId)
	str := "./images/upload" + s + ".png"
	f, _ := os.Create(str)
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	w.WriteHeader(http.StatusOK)
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		user := Models.User{Login: "", Email: "", Password: ""}
		u := &user
		result, err := json.Marshal(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(result)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	w.WriteHeader(http.StatusOK)
	user := Models.User{Login: "", Email: "", Password: ""}
	u := &user
	result, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(result)
}
