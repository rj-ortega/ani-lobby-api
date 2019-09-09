package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User model
type User struct {
	ID        string    `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `json:"username"`
	Animes    []*Anime  `gorm:"many2many:user_languages;"`
}

func getAllUsers(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"message": users,
	})
}

func getUser(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	var user User
	db.Where("id = ?", c.Param("id")).First(&user)
	if id := user.ID; id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func createUser(c *gin.Context) {
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

func deleteUser(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	id := c.Param("id")
	var user = User{ID: id}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}

func updateUser(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	var body User
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	db.Model(&User{}).Where("id = ?", id).Updates(body)
	result := User{ID: id}
	db.Find(&result)
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func updateAnimeRelationship(c *gin.Context) {
	db := c.MustGet(DBName).(*gorm.DB)
	var user User
	var anime Anime
	user.ID = c.Param("id")
	anime.ID = c.Param("anime_id")
	db.Model(&user).Association("Anime").Find(&anime)
	if anime.Title == "" {
		// si la db no tiene el anime
		// crea el anime con la informacion del body o del api??
	}
	db.Model(&user).Association("Anime").Append(anime)
	c.JSON(http.StatusOK, gin.H{
		"message": "anime associated to user",
	})
}
