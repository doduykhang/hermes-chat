package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	envFile = map[string]string {
		"DEV": "./env/dev.json",
		"PROD": "./env/prod.json",
	}
)

type Redis struct {
	Host string `json:"HOST"`	
	Port string `json:"PORT"`	
	Username string `json:"USERNAME"`
	Password string `json:"PASSWORD"`
}

type Env struct{
	Port string `json:"PORT"`
	Redis Redis `json:"REDIS"`
}

var (
	ENV Env
)

func init() {
	envFilePath := envFile["DEV"]
	if file := envFile[os.Getenv("ENV")]; file != ""  {
		envFilePath = file
	}
	bytes, err := os.ReadFile(envFilePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &ENV)
	if err != nil {
		log.Fatal(err)
	}
}

func GetEnv () *Env {
	return &ENV
}


