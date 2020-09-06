package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DBConf struct {
	URL                string
	MaxOpenConnections int
	MaxIdleConnections int
	DbEngine           string
}

type JwtConf struct {
	timeout    int32
	maxRefresh int32
}

type Config struct {
	DB  DBConf
	JWT JwtConf
}

//Configuration Global configuration variable
var Configuration *Config

//InitConfiguration load configuration file
func InitConfiguration() {
	fmt.Println("Init Config")
	file, _ := os.Open("./config/test-config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)

	err := decoder.Decode(&Configuration)
	fmt.Println(Configuration)
	if err != nil {
		panic("config file not found")
	}
}
