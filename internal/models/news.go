package models

import "time"

type News struct {
	Id        string    `json:"id" form:"id"`
	Title     string    `json:"title" form:"title"`
	Body      string    `json:"body" form:"body"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	Position  string    `json:"position"  form:"position"`
}
