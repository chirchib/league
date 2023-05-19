package router

import "github.com/gin-gonic/gin"

func routes(g *gin.Engine, d *routingDependencies) {
	// routes
	v1Group := g.Group("v1")

	// clubs
	v1Group.GET("/clubs", d.clubsHandler.GetClubs)
	v1Group.GET("/clubs/:club_id", d.clubsHandler.GetClub)
	v1Group.GET("/clubs/:club_id/players", d.clubsHandler.GetClubPlayers)
	v1Group.GET("/clubs/:club_id/titles", d.clubsHandler.GetClubTitles)
	v1Group.GET("/clubs/:club_id/games", d.clubsHandler.GetClubGames)

	// seasons
	v1Group.GET("/seasons", d.seasonsHandler.GetSeasons)
	v1Group.GET("/seasons/:season_id", d.seasonsHandler.GetSeason)

	// league_tables
	v1Group.GET("/seasons/:season_id/league_tables", d.tablesHandler.GetTable)
	v1Group.GET("/seasons/:season_id/league_tables/statistics", d.tablesHandler.GetStat)
	v1Group.GET("/seasons/:season_id/league_tables/games", d.tablesHandler.GetGames)
	v1Group.GET("/seasons/:season_id/league_tables/:club_id/stats", d.tablesHandler.GetClubStats)

	// players
	v1Group.GET("/players", d.playersHandler.GetPlayers)
	v1Group.GET("/players/:player_id", d.playersHandler.GetPlayer)
	v1Group.GET("/players/:player_id/stats", d.playersHandler.GetPlayerStats)
}
