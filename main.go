package main

import "github.com/gin-gonic/gin"

// Anime model
type Anime struct {
	Name           string   `json:"name"`
	Genres         []string `json:"genres"`
	MyAnimeListURL string   `json:"my_anime_list_url"`
	Episodes       []string `json:"episodes"`
}

// User model
type User struct {
	Name string `json:"name"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
