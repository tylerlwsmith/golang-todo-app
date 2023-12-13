package models

import "time"

type pageData interface{}

type Page struct {
	Title    string
	Content  string
	PageData pageData
}

type Task struct {
	Id          int
	Description string
	Completed   bool
	CreatedAt   time.Time
	ModifiedAt  time.Time
}
