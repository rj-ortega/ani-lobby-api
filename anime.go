package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

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
	Users     []*User   `gorm:"many2many:user_animes;"`
}

// Genre is the genre information from season anime
type Genre struct {
	MalID int    `json:"mal_id"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func getAllAnimes(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	var animes []Anime
	db.Find(&animes)
	c.JSON(http.StatusOK, gin.H{
		"message": animes,
	})
}

func getAnime(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	var anime Anime
	db.Where("id = ?", c.Param("id")).First(&anime)
	if id := anime.ID; id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No anime found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": anime,
	})
}

func createAnime(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	var anime Anime
	if err := c.ShouldBind(&anime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if animeID := anime.ID; animeID == "" {
		animeID = uuid.New().String()
	}
	db.Create(&anime)
	c.JSON(http.StatusOK, gin.H{
		"message": anime,
	})
}

func deleteAnime(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	id := c.Param("id")
	var anime = Anime{ID: id}
	db.Delete(&anime)
	c.JSON(http.StatusOK, gin.H{
		"message": "Anime deleted",
	})
}

func updateAnime(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	var body Anime
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	db.Model(&Anime{}).Where("id = ?", id).Updates(body)
	result := Anime{ID: id}
	db.Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
