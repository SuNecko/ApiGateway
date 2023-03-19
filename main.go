package main

import (
	"Baseball_Api/handlers"
	"Baseball_Api/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

func main() {
	r := gin.Default()

	group := r.Group("api/v1")

	//Teams Handlres
	group.GET("/teams", handlers.GetTeams)
	group.GET("/teams/:id", handlers.GetTeamById)
	group.GET("/statistics/:team/:season", handlers.GetTeamStatistics)

	//Game Handlres
	group.GET("/games/today", handlers.GetGameToDay)
	group.GET("/games/:season", handlers.GetGamesBySeason)
	group.GET("/games/:season/:date", handlers.GetGamesByDate)
	group.GET("/teams/games/:season/:team", handlers.GetTeamGamesBySeason)
	group.GET("/teams/games/:season/:team/:date", handlers.GetTeamGamesByDate)

	//Game vs two teams Handlres
	group.GET("/games/vs/:season/:h2h", handlers.GetGameVs)

	//Standings Handler
	group.GET("/standings/:season", handlers.GetStandingsBySeason)
	group.GET("/standings/:season/:team", handlers.GetTeamStanding)

	var mutex sync.Mutex
	errch := make(chan error)
	go tools.RefreshRequestsData(errch, &mutex)
	go tools.ErrorHandler(errch)
	go tools.GarbageCollector()

	r.Run("localhost: 8080")
	fmt.Print("Corriendo")
}
