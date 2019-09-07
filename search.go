package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Search returns a list of anime
type Search struct {
	Results []Results `json:"results"`
}

// Results returns a list of anime that match query
type Results struct {
	MalID    int     `json:"mal_id"`
	Title    string  `json:"title"`
	URL      string  `json:"url"`
	ImageURL string  `json:"image_url"`
	Score    float64 `json:"score"`
	Episodes uint    `json:"episodes"`
	Synopsis string  `json:"synopsis" gorm:"type:text"`
}

func searchForAnime(c *gin.Context) {
	query := c.Query("search")
	queryURL := fmt.Sprintf("%s/search/anime/?q=%s&limit=10", AnimeAPIURL, query)
	resp, err := http.Get(queryURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No anime found"})
		return
	}

	var search Search
	json.NewDecoder(resp.Body).Decode(&search)

	c.JSON(http.StatusOK, gin.H{
		"message": search,
	})
}
