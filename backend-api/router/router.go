package router

import (
	"backend/db"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	// 同一オリジンポリシーの設定（ミドルウェア）
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("FRONTEND_HOST")},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	router.GET("/ping", ping)
	router.GET("/api/v1/articles", articles)
	router.GET("/api/v1/articles/:id", article)

	Router = router
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func articles(c *gin.Context) {
	res, err := db.GetAllArticles()
	fmt.Fprintf(os.Stdout, "res: %v\n", res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func article(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	res, err := db.GetArticle(id)
	fmt.Fprintf(os.Stdout, "res: %v\n", res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
