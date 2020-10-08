package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"./Handlers"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/signup", Handlers.SignupPage)
	router.HandleFunc("/login", Handlers.LoginPage)
	router.HandleFunc("/me", Handlers.IsMe)
	router.HandleFunc("/whois", Handlers.Whois)
	router.HandleFunc("/chengelogin", Handlers.Chengelogin)
	router.HandleFunc("/chengepass", Handlers.Chengepass)
	router.HandleFunc("/chengeavatar", Handlers.Chengeavatar)
	router.HandleFunc("/logout", Handlers.LogoutPage)
	router.HandleFunc("/", Handlers.MainPage)
	router.Use(Handlers.CORSMiddleware())
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../public"))))
	err := http.ListenAndServe(":8000", router)
	fmt.Println("Listen 8000 port")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}