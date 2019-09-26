package db

import (
	"backend/models"
	"fmt"
	"os"
)

func GetAllArticles() (articles []models.Article, err error) {
	stmt, err := Conn.Query("select * from articles")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	defer stmt.Close()
	for stmt.Next() {
		var id int
		var title, description string
		if err := stmt.Scan(&id, &title, &description); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
		article := models.Article{
			Id:          id,
			Title:       title,
			Description: description,
		}
		articles = append(articles, article)
	}
	return
}

func GetArticle(paramId int) (article models.Article, err error) {
	stmt := Conn.QueryRow("select * from articles where id = ? limit 1", paramId)
	var id int
	var title, description string
	stmt.Scan(&id, &title, &description)
	article = models.Article{
		Id:          id,
		Title:       title,
		Description: description,
	}

	return
}
