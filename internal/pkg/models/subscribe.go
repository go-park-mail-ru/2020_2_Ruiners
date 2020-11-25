package models

type Subscribe struct {
	Id int `json:"id"`
	Body string `json:"body"`
	Date int64
}

type Feed []Subscribe