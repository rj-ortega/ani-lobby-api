package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DBName refers to name of database
const DBName = "ani_lobby"
const localConnection = "host=127.0.0.1 port=5432 user=Rj dbname=ani_lobby sslmode=disable"

// AnimeAPIURL is the url for the api supplying information
const AnimeAPIURL = "https://api.jikan.moe/v3"

// Anime model
type Anime struct {
	ID        string    `json:"id" gorm:"primary_key"`
	MalID     int       `json:"mal_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	ImageURL  string    `json:"image_url"`
	Score     float64   `json:"score"`
	Episodes  uint      `json:"episodes"`
	Synopsis  string    `json:"synopsis" gorm:"type:text"`
	Genres    []Genre   `json:"genres"`
}

// Genre is the genre information from season anime
type Genre struct {
	MalID int    `json:"mal_id"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func addDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		dbURL := os.Getenv("DATABASE_URL")
		if dbURL == "" {
			dbURL = localConnection
		}
		db, err := gorm.Open("postgres", dbURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Set(DBName, db)
		db.AutoMigrate(&User{})
		db.AutoMigrate(&Anime{})
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
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Anime Lobby")
	})
	user := r.Group("/api/v1/users")
	{
		user.GET("", getAllUsers)
		user.GET("/:id", getUser)
		user.POST("", createUser)
		user.DELETE("/:id", deleteUser)
		user.PATCH("/:id", updateUser)
	}
	anime := r.Group("/api/v1/animes")
	{
		anime.GET("", getAllAnimes)
		anime.GET("/:id", getAnime)
		anime.POST("", createAnime)
		anime.DELETE("/:id", deleteAnime)
		anime.PATCH("/:id", updateAnime)
	}
	r.GET("/api/v1/seasons", getSeasonalAnimes)
	r.GET("api/v1/search", searchForAnime)

	r.Run(fmt.Sprintf(":%s", port))
}
