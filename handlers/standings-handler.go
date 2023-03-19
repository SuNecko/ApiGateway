package handlers

import (
	"Baseball_Api/tools"
	"github.com/gin-gonic/gin"
)

func GetStandingsBySeason(c *gin.Context) {
	season := c.Param("season")

	url := "https://v1.baseball.api-sports.io/standings?league=9&season=" + season

	tools.Response(url, c)
}

func GetTeamStanding(c *gin.Context) {
	season := c.Param("season")
	team := c.Param("team")

	url := "https://v1.baseball.api-sports.io/standings?league=9&season=" + season + "&team=" + team

	tools.Response(url, c)
}
