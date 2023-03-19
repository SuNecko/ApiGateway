package handlers

import (
	"Baseball_Api/tools"
	"github.com/gin-gonic/gin"
)

func GetGameVs(c *gin.Context) {
	season := c.Param("season")
	h2h := c.Param("h2h")

	url := "https://v1.baseball.api-sports.io/games/h2h?season=" + season + "&h2h=" + h2h

	tools.Response(url, c)
}
