package models

import "time"

type Post struct {
	Title   string
	Content string
	Author  string
	Date    time.Time
}
