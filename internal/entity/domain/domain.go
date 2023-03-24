package domain

import "time"

type Team struct {
	ID       string
	Name     string
	Manager  string
	Stadium  string
	LogoTeam string
}

type Player struct {
	ID       string
	Name     string
	Position string
	Birthday time.Time
	Nations  string
	TeamID   string
}

type Game struct {
	ID        string
	Date      time.Time
	HouseTeam string
	Result    string
	Stadium   string
}

type Season struct {
	ID          string
	StartDate   time.Time
	EndDate     time.Time
	TeamsNumber int
}

type LeagueTable struct {
	ID        string
	TeamName  string
	Win       int
	Draw      int
	Lose      int
	DiffGoals int
	Scores    int
}

type Statistic struct {
	ID          string
	PlayerName  string
	Goals       int
	GoalPass    int
	YellowCards int
	RedCarts    int
}
