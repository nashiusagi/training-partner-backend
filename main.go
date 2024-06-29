package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

func main() {
	db, err := gorm.Open("sqlite3", "resources/post.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Post{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/posts", func(c *gin.Context) {
		var posts []Post
		if err := db.Find(&posts).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		} else {
			db.Find(&posts)
			c.JSON(http.StatusOK, posts)
		}
	})

	r.Run()
}
