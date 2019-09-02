package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const DBName = "ani_lobby"
const localConnection = "host=127.0.0.1 port=5432 user=Rj dbname=ani_lobby sslmode=disable"

// Anime model
type Anime struct {
	Name           string   `json:"name"`
	Genres         []string `json:"genres"`
	MyAnimeListURL string   `json:"my_anime_list_url"`
	Episodes       []string `json:"episodes"`
}

func addDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbURL := os.Getenv("DATABASE_URL")
		if dbURL == "" {
			dbURL = localConnection
		}
		db, err := gorm.Open("postgres", dbURL)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Set(DBName, db)
		db.AutoMigrate(&User{})
		c.Next()
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := gin.Default()
	r.Use(addDB())
	r.GET("/users", getAllUsers)
	r.GET("/users/:id", getUser)
	// r.POST("/users", getAllUsers)
	// r.DELETE("/users", getAllUsers)
	// r.PATCH("/users", getAllUsers)
	r.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080
}
