package config

import (
	"encoding/json"
	"os"
)

type DBConf struct {
	url string
}

type JwtConf struct {
	timeout    int32
	maxRefresh int32
}

type Config struct {
	DB  DBConf
	JWT JwtConf
}

func InitConfiguration() {
	file, _ := os.Open("./config/test-config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Config{}

	err := decoder.Decode(&configuration)
	if err != nil {
		panic("config file not found")
	}
}
