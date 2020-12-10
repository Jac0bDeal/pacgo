package pacgo

import (
	"encoding/json"
	"os"
)

type Config struct {
	PlayerSprite string `json:"player_sprite"`
	GhostSprite  string `json:"ghost_sprite"`
	WallSprite   string `json:"wall_sprite"`
	DotSprite    string `json:"dot_sprite"`
	PillSprite   string `json:"pill_sprite"`
	DeathSprite  string `json:"death_sprite"`
	SpaceSprite  string `json:"space_sprite"`
	UseEmoji     bool   `json:"use_emoji"`
}

func loadConfig(filepath string) (*Config, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
