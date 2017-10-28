package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Server     string
	MongoDbURL string
}

var appConfig *Configuration

var configPath = "config/config.json"

func GetConfig() Configuration {
	if appConfig == nil {
		appConfig = &Configuration{}

		file, err := os.Open(configPath)
		defer file.Close()

		if err != nil {
			log.Fatalf("[loadConfig]: %s\n", err)
		}

		decoder := json.NewDecoder(file)

		if err = decoder.Decode(appConfig); err != nil {
			log.Fatalf("[logAppConfig]: %s\n", err)
		}

		return *appConfig
	}

	return *appConfig
}
