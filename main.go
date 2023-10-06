package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	steamID := flag.String("id", "", "Steam ID")
	matchCount := flag.Int("count", 1000, "Number of matches to retrieve")

	flag.Parse()

	if *steamID == "" {
		fmt.Println("Please provide a steam ID")
		return
	}

	api := AOE2NET{
		MatchAPI: "https://aoe2.net/api/player/matches?game=aoe2de&steam_id=%s&count=%d&start=%d",
		CivAPI:   "https://aoe2.net/api/strings?game=aoe2de&language=en",
	}

	civilizations, err := api.GetCivilazations()
	if err != nil {
		panic(err)
	}

	matches, err := api.GetAllMatches(*steamID, *matchCount)
	if err != nil {
		fmt.Println("error getting matches but attempting to continue", err)
	}

	playstyle := CalculatePlaystyle(*steamID, civilizations, matches)

	fmt.Printf("%+v\n", playstyle)
	fmt.Printf("Found %d matches, %f", len(matches), playstyle.Versatility())

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	pc := widgets.NewPieChart()
	pc.Title = "Race Playstyle"
	pc.SetRect(5, 5, 70, 36)
	pc.Data = []float64{
		playstyle.Archers / playstyle.Versatility() * 100,
		playstyle.Cavalry / playstyle.Versatility() * 100,
		playstyle.CavalryArchers / playstyle.Versatility() * 100,
		playstyle.Infantry / playstyle.Versatility() * 100,
		playstyle.Siege / playstyle.Versatility() * 100,
		playstyle.Monks / playstyle.Versatility() * 100,
		playstyle.Water / playstyle.Versatility() * 100,
	}
	pc.Colors = []ui.Color{130, 5, 1, ui.ColorWhite, 8, 3, 12}
	pc.AngleOffset = -.5 * math.Pi
	pc.LabelFormatter = func(i int, v float64) string {
		var label string

		switch i {
		case 0:
			label = "Archers"
		case 1:
			label = "Cavalry"
		case 2:
			label = "Cavalry Archers"
		case 3:
			label = "Infantry"
		case 4:
			label = "Siege"
		case 5:
			label = "Monks"
		case 6:
			label = "Water"
		}

		return fmt.Sprintf("%s: %.2f%%", label, v)
	}

	ui.Render(pc)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}

type Civilization struct {
	ID        int    `json:"id"`
	Name      string `json:"string"`
	Playstyle Playstyle
}

type Match struct {
	MatchID           string      `json:"match_id"`
	LobbyID           interface{} `json:"lobby_id"`
	MatchUUID         string      `json:"match_uuid"`
	Version           interface{} `json:"version"`
	Name              string      `json:"name"`
	NumPlayers        int         `json:"num_players"`
	NumSlots          int         `json:"num_slots"`
	AverageRating     interface{} `json:"average_rating"`
	Cheats            bool        `json:"cheats"`
	FullTechTree      bool        `json:"full_tech_tree"`
	EndingAge         int         `json:"ending_age"`
	Expansion         interface{} `json:"expansion"`
	GameType          int         `json:"game_type"`
	HasCustomContent  interface{} `json:"has_custom_content"`
	HasPassword       interface{} `json:"has_password"`
	LockSpeed         bool        `json:"lock_speed"`
	LockTeams         bool        `json:"lock_teams"`
	MapSize           int         `json:"map_size"`
	MapType           int         `json:"map_type"`
	Pop               int         `json:"pop"`
	Ranked            bool        `json:"ranked"`
	LeaderboardID     int         `json:"leaderboard_id"`
	RatingType        int         `json:"rating_type"`
	Resources         int         `json:"resources"`
	Rms               interface{} `json:"rms"`
	Scenario          interface{} `json:"scenario"`
	Server            interface{} `json:"server"`
	SharedExploration bool        `json:"shared_exploration"`
	Speed             int         `json:"speed"`
	StartingAge       int         `json:"starting_age"`
	TeamTogether      bool        `json:"team_together"`
	TeamPositions     bool        `json:"team_positions"`
	TreatyLength      int         `json:"treaty_length"`
	Turbo             bool        `json:"turbo"`
	Victory           int         `json:"victory"`
	VictoryTime       int         `json:"victory_time"`
	Visibility        int         `json:"visibility"`
	Opened            int         `json:"opened"`
	Started           int         `json:"started"`
	Finished          int         `json:"finished"`
	Players           []struct {
		ProfileID    int         `json:"profile_id"`
		SteamID      string      `json:"steam_id"`
		Name         string      `json:"name"`
		Clan         interface{} `json:"clan"`
		Country      string      `json:"country"`
		Slot         int         `json:"slot"`
		SlotType     int         `json:"slot_type"`
		Rating       int         `json:"rating"`
		RatingChange interface{} `json:"rating_change"`
		Games        interface{} `json:"games"`
		Wins         interface{} `json:"wins"`
		Streak       interface{} `json:"streak"`
		Drops        interface{} `json:"drops"`
		Color        int         `json:"color"`
		Team         int         `json:"team"`
		Civ          int         `json:"civ"`
		CivAlpha     int         `json:"civ_alpha"`
		Won          bool        `json:"won"`
	} `json:"players"`
}

type AOE2NET struct {
	MatchAPI string
	CivAPI   string
}

func (a AOE2NET) GetCivilazations() ([]*Civilization, error) {
	var civilizations []*Civilization

	r, err := http.Get(a.CivAPI)
	if err != nil {
		return civilizations, err
	}

	response := struct {
		Civilizations []*Civilization `json:"civ"`
	}{}

	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return civilizations, err
	}

	for _, civilization := range response.Civilizations {
		playstyle, err := GetPlaystyle(civilization.Name)
		if err != nil {
			fmt.Println("failed to retrieve playstyle", err)
			continue
		}
		civilization.Playstyle = playstyle
	}

	return response.Civilizations, nil
}

func (a AOE2NET) GetMatches(steamID string, count int, startIndex int) ([]Match, error) {
	var matches []Match

	matchAPI := "https://aoe2.net/api/player/matches?game=aoe2de&steam_id=%s&count=%d&start=%d"

	r, err := http.Get(fmt.Sprintf(matchAPI, steamID, count, startIndex))
	if err != nil {
		return matches, err
	}

	err = json.NewDecoder(r.Body).Decode(&matches)
	if err != nil {
		return matches, err
	}

	return matches, nil
}

func (a AOE2NET) GetAllMatches(steamID string, amount int) ([]Match, error) {
	var allMatches []Match

	var matchesRetrieved int

	for matchesRetrieved < amount {
		batchAmount := amount - matchesRetrieved
		if batchAmount > 1000 {
			batchAmount = 1000
		}
		matches, err := a.GetMatches(steamID, batchAmount, matchesRetrieved)
		if err != nil {
			return allMatches, err
		}

		if len(matches) == 0 {
			break
		}

		matchesRetrieved += len(matches)
		allMatches = append(allMatches, matches...)
	}

	return allMatches, nil
}

func CalculatePlaystyle(steamID string, civilizations []*Civilization, matches []Match) Playstyle {
	playerPlaystyle := Playstyle{}

	for _, match := range matches {
		for _, player := range match.Players {
			if player.SteamID == steamID {
				for _, civilization := range civilizations {
					if player.Civ == civilization.ID {
						playerPlaystyle.Archers += civilization.Playstyle.Archers
						playerPlaystyle.Cavalry += civilization.Playstyle.Cavalry
						playerPlaystyle.CavalryArchers += civilization.Playstyle.CavalryArchers
						playerPlaystyle.Infantry += civilization.Playstyle.Infantry
						playerPlaystyle.Siege += civilization.Playstyle.Siege
						playerPlaystyle.Monks += civilization.Playstyle.Monks
						playerPlaystyle.Water += civilization.Playstyle.Water
					}
				}
			}
		}
	}

	return playerPlaystyle
}
