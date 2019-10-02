package router

import (
	"backend/db"
	"backend/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"cloud.google.com/go/datastore"
	"time"
	"github.com/golang/glog"
)

var Router *gin.Engine
var ProjectID = os.Getenv("PROJECT_ID")

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

	router.GET("/api/v1/posts", GetPosts)
	router.POST("/api/v1/post", NewPost)
	router.GET("/api/v1/post/:slug", GetPost)
	router.POST("/api/v1/post/:slug", UpdatePost)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not Found"})
	})

	Router = router
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func articles(c *gin.Context) {
	res, err := db.GetAllArticles()
	glog.Infof("res: %v\n", res)
	if err != nil {
		glog.Errorf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func article(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		glog.Errorf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	res, err := db.GetArticle(id)
	glog.Infof("Error: %v\n", err)
	if err != nil {
		glog.Errorf("Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}


func GetPosts(c *gin.Context) {
	client, err := datastore.NewClient(c, ProjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"InternalServerError": err})
		glog.Errorf("Faild to connect datastore (reason: %v)\n",err)
		return
	}
	defer client.Close()

	var posts []models.Post

	_, err = client.GetAll(c, datastore.NewQuery("Post").Limit(20), &posts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"InternalServerError": err})
		glog.Errorf("Faild to get all posts (reason: %v)\n",err)
		return
	}

	glog.Infof( "res: %v\n", posts)

	c.JSON(http.StatusOK, posts)
}


func GetPost(c *gin.Context) {
	client, err := datastore.NewClient(c, ProjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"InternalServerError": err})
		glog.Errorf("Faild to connect datastore (reason: %v)\n",err)
		return
	}
	defer client.Close()

	var post models.Post

	key := datastore.NameKey("Post", c.Param("slug"), nil)

	err = client.Get(c, key, &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"InternalServerError": "hogehoge"})
		glog.Errorf("Faild to get specifeid post (reason: %v)\n",err)
		return
	}

	glog.Infof("res: %v\n", post)

	c.JSON(http.StatusOK, post)
}


func NewPost(c *gin.Context) {
	client, err := datastore.NewClient(c, ProjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"InternalServerError": err})
		glog.Errorf("Faild to connect to datastore reason: %v\n", err)
		return
	}
	defer client.Close()

	var post models.Post
	c.BindJSON(&post)
	glog.Infof("POST_DATA: %v\n", post)
	c.JSON(http.StatusOK, post)


	if post.Slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Slug is required"})
		glog.Errorln("Error : Slug is required")
		return
	}

	if post.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Title is required"})
		glog.Errorln("Error : Title is required")
		return
	}

	//u, err := uuid.NewRandom()
	//if err != nil {
	//	log.Fatalf("Faild to create uuid (reason: %v)", err)
	//	return
	//}
	//post.Id = u.String()

	post.Created = time.Now()
	post.Updated = time.Now()

	key := datastore.NameKey("Post", post.Slug, nil)

	res, err := client.Put(c, key, &post);
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"InternalServerError": err})
		glog.Errorf("Failed to save post: %v", err)
		return
	}

	c.JSON(http.StatusOK, res)
}


func UpdatePost(c *gin.Context) {
	client, err := datastore.NewClient(c, ProjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"InternalServerError": err})
		glog.Errorf("Faild to connect to datastore reason: %v\n", err)
		return
	}
	defer client.Close()

	var post models.Post

	key := datastore.NameKey("Post", c.Param("slug"), nil)

	err = client.Get(c, key, &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"InternalServerError": err})
		glog.Errorf("Faild to get all posts (reason: %v)\n",err)
		return
	}

	c.BindJSON(&post)
	glog.Infof("POST_DATA: %v\n", post)

	post.Updated = time.Now()

	if post.Slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Slug is required"})
		glog.Errorln("Error : Slug is required")
		return
	}

	if post.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Title is required"})
		glog.Errorln("Error : Title is required")
		return
	}

	err = client.Delete(c, key)

	key = datastore.NameKey("Post", post.Slug, nil)

	res, err := client.Put(c, key, &post);
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"InternalServerError": err})
		glog.Errorf("Failed to save post: %v", err)
		return
	}

	c.JSON(http.StatusOK, res)
}
