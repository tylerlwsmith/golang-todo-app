package models

import "time"

type pageData interface{}

type Page struct {
	Title    string
	Content  string
	PageData pageData
}

type Todo struct {
	Id          int
	Description string
	Completed   bool
	CreatedAt   time.Time
	ModifiedAt  time.Time
}
