package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DBName refers to name of database
const DBName = "ani_lobby"
const localConnection = "host=127.0.0.1 port=5432 user=Rj dbname=ani_lobby sslmode=disable"

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
	users := r.Group("/api/v1/users")
	{
		users.GET("", getAllUsers)
		users.GET("/:id", getUser)
		users.POST("", createUser)
		users.DELETE("/:id", deleteUser)
		users.PATCH("/:id", updateUser)
	}
	anime := r.Group("/api/v1/animes")
	{
		anime.GET("", getAllAnimes)
		anime.GET("/:id", getAnime)
		anime.POST("", createAnime)
		// anime.DELETE("", deleteAnime)
		// anime.PATCH("", updateAnime)
	}
	r.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080
}
