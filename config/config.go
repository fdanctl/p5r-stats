package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerPort string `json:"server_port"`
	DataFile   string `json:"data_file"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
