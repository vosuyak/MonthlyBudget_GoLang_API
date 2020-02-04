package core

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Configuration
type Configuration struct {
	Port          string `json:"port"`
	LogFilename   string `json:"logFilename"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`

	MgAddrs      string `json:"mgAddrs"`
	MgDbName     string `json:"mgDbName"`
	MgDbUsername string `json:"mgDbUsername"`
	MgDbPassword string `json:"mgDbPassword"`

	JwtSecretPassword string `json:"jwtSecretPassword"`
	Issuer            string `json:"issuer"`
}

// AppConfig holds the configuration values from config.json file
var AppConfig Configuration

// LoadAppConfig reads config.json and decode into AppConfig
func LoadAppConfig() {
	file, err := os.Open("core/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = Configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}

// LoadEnv - load enviorment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env := os.Getenv("APP_ENV")
	if "" == env {
		log.Fatal("No variabels in file")
	} else {
		log.Println("using: ","./core/.env." + env)
		err = godotenv.Load("./core/.env." + env)
		if err != nil {
			log.Fatal("Error loading .env." + env + " file access might be restricted")
		}
	}
}
