package config

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Server                 string
	ClientsServiceURL      string
	DestinationsServiceURL string
	SourcesServiceURL      string
}

var appConfig = &Configuration{}

var configPath = "config/config.json"

func GetConfig() Configuration {
	file, err := os.Open(configPath)
	defer file.Close()

	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}

	if err = json.NewDecoder(file).Decode(appConfig); err != nil {
		log.Fatalf("[logAppConfig]: %s\n", err)
	}

	return *appConfig
}
