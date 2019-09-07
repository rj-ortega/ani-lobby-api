package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Season represents anime response for the season
type Season struct {
	SeasonName string  `json:"season_name"`
	SeasonYear int     `json:"season_year"`
	Anime      []Anime `json:"anime"`
}

func getSeasonalAnimes(c *gin.Context) {
	s := c.Query("season")
	y := c.Query("year")
	resp, err := http.Get(fmt.Sprintf("%s/season/%s/%s", AnimeAPIURL, y, s))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No anime found"})
		return
	}

	var season Season
	json.NewDecoder(resp.Body).Decode(&season)

	c.JSON(http.StatusOK, gin.H{
		"message": season,
	})
}
