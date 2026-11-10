package statistics

import (
	"encoding/json"
	"os"
)

type Stats struct {
	Player1Wins  int `json:"player1_wins"`
	Player2Wins  int `json:"player2_wins"`
	ComputerWins int `json:"computer_wins"`
	Draws        int `json:"draws"`
	TotalGames   int `json:"total_games"`
}

func Load(filename string) (Stats, error) {
	var stats Stats

	data, err := os.ReadFile(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return stats, nil
		}
		return stats, err
	}

	err = json.Unmarshal(data, &stats)

	return stats, err
}

func Save(filename string, stats Stats) error {
	data, err := json.MarshalIndent(stats, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (s *Stats) Record(result string, mode string) {
	s.TotalGames++

	if result == "X" {
		s.Player1Wins++
	} else if result == "O" && mode == "computer" {
		s.ComputerWins++
	} else if result == "O" {
		s.Player2Wins++
	} else {
		s.Draws++
	}
}
