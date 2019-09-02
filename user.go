package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User model
type User struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Username  string     `json:"username"`
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
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

// func createUser(c *gin.Context) {
// 	db := c.MustGet(DBName).(*gorm.DB)

// }
