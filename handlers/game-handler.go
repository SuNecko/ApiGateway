package handlers

import (
	"Baseball_Api/tools"
	"github.com/gin-gonic/gin"
	"time"
)

func GetGameToDay(c *gin.Context) {
	t := time.Now()
	t.String()
	date := t.Format("2006-01-02")

	url := "https://v1.baseball.api-sports.io/games?league=9&season=2022&date=" + date

	tools.Response(url, c)
}

func GetGamesBySeason(c *gin.Context) {
	season := c.Param("season")

	url := "https://v1.baseball.api-sports.io/games?league=9&season=" + season

	tools.Response(url, c)
}

func GetGamesByDate(c *gin.Context) {
	season := c.Param("season")
	date := c.Param("date")

	url := "https://v1.baseball.api-sports.io/games?league=9&season=" + season + "&date=" + date

	tools.Response(url, c)
}

func GetTeamGamesBySeason(c *gin.Context) {
	season := c.Param("season")
	team := c.Param("team")

	url := "https://v1.baseball.api-sports.io/games?league=9&season=" + season + "&team=" + team

	tools.Response(url, c)
}

func GetTeamGamesByDate(c *gin.Context) {
	season := c.Param("season")
	date := c.Param("date")
	team := c.Param("team")

	url := "https://v1.baseball.api-sports.io/games?league=9&season=" + season + "&date=" + date + "&team=" + team

	tools.Response(url, c)
}
