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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ImageURL  string    `jsong:"image_url"`
	Score     float64   `json:"score"`
	Episodes  uint      `json:"episodes"`
	Synopsis  string    `json:"synopsis" gorm:"type:text"`
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
	c.JSON(http.StatusOK, gin.H{
		"message": anime,
	})
}

func createAnime(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uuid.New().String()
	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

// func updateAnime(c *gin.Context) {
// 	db := c.MustGet(DBName).(*gorm.DB)
// 	 := c.Query("")
// 	var anime = Anime{}
// 	db.Model(&anime).Update("", )
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": anime,
// 	})
// }

// func deleteAnime(c *gin.Context) {
// 	db := c.MustGet(DBName).(*gorm.DB)
// 	id := c.Query("id")
// 	var anime = Anime{ID: id}
// 	db.Delete(&anime)
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": anime.Name + "deleted",
// 	})
// }
