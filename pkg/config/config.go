package config

import (
	"encoding/json"
	"os"

	"github.com/hinotora/go-auth-service/pkg/logger"
)

type Config struct {
	Application struct {
		Url  string `json:"url"`
		Port int8   `json:"port"`
		Env  string `json:"environment"`
	} `json:"app"`
	Database struct {
		Host     string `json:"host"`
		Port     int16  `json:"port"`
		Database string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"db"`
	Auth struct {
		Jwt_secret_key   string `json:"jwt-secret-key"`
		Jwt_time_to_live int16  `json:"jwt-time-to-live"`
	} `json:"auth"`
}

var Instance *Config = nil

var ConfigFileName string = "config.json"

func Load() error {
	if Instance != nil {
		return nil
	}

	logger.Logger.Println("Using config " + ConfigFileName)

	jsonFile, err := os.ReadFile(ConfigFileName)

	if err != nil {
		return err
	}

	config := &Config{}

	err = json.Unmarshal(jsonFile, config)

	if err != nil {
		return err
	}

	Instance = config

	logger.Logger.Println("Config loaded succesfully")

	return nil
}
