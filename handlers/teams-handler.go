package handlers

import (
	"Baseball_Api/tools"
	"github.com/gin-gonic/gin"
)

func GetTeams(c *gin.Context) {
	url := "https://v1.baseball.api-sports.io/teams?country_id=7"

	tools.Response(url, c)
}

func GetTeamById(c *gin.Context) {
	id := c.Param("id")

	url := "https://v1.baseball.api-sports.io/teams?id=" + id

	tools.Response(url, c)
}

func GetTeamStatistics(c *gin.Context) {
	team := c.Param("team")
	season := c.Param("season")

	url := "https://v1.baseball.api-sports.io/teams/statistics?league=9&team=" + team + "&season=" + season + ""

	tools.Response(url, c)

}
