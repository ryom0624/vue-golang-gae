package models

type Article struct {
	Id          int
	Title       string
	Description string
	Articles    []Article
}
