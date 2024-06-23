package model

type Todo struct {
	Id     int16  `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
