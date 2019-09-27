package models

import "time"

type Post struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	Created  time.Time
	Updated  time.Time
}
