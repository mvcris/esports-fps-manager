package backend

import (
	"encoding/json"
	"esports-fps-manager/backend/entities"
	"fmt"
	"log"
	"os"
)

type Game struct {
	GameName string
}

func (g *Game) LoadPlayers() {
	playersFile, err := os.ReadFile("data/player.json")
	if err != nil {
		log.Fatalf("Error reading players file: %v", err)
	}

	var players []entities.Player
	json.Unmarshal(playersFile, &players)
	fmt.Println(players)
}

func (g *Game) LoadTeams() {
	teamsFile, err := os.ReadFile("data/team.json")
	if err != nil {
		log.Fatalf("Error reading teams file: %v", err)
	}

	var teams []entities.Team
	json.Unmarshal(teamsFile, &teams)
	fmt.Println(teams)
}

func (g *Game) LoadTeamPlayers() {
	teamPlayersFile, err := os.ReadFile("data/team_player.json")
	if err != nil {
		log.Fatalf("Error reading team players file: %v", err)
	}

	var teamPlayers []entities.TeamPlayer
	json.Unmarshal(teamPlayersFile, &teamPlayers)
	fmt.Println(teamPlayers)

}

func (g *Game) CreateGame() {}

func LoadGame() {}
