package config

import (
	"encoding/json"
	"os"

	"github.com/mateicheles1/golang-crud/logs"
)

func NewConfig(configFilePath string) Config {
	file, err := os.Open(configFilePath)

	if err != nil {
		logs.Logger.Fatal().Msgf("Error opening config file path: %s", err)
	}

	defer file.Close()

	var config Config

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		logs.Logger.Fatal().Msgf("Error unmarshaling json into config: %s", err)
	}

	return config

}
