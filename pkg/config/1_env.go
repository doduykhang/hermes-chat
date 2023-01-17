package config

import (
	"encoding/json"
	"log"
	"os"
)


type Redis struct {
	Host string `json:"HOST"`	
	Port string `json:"PORT"`	
	Username string `json:"USERNAME"`
	Password string `json:"PASSWORD"`
}

type Env struct{
	Port string `json:"PORT"`
	JwtSecrect string `json:"JWT_SCERECT"`
	Redis Redis `json:"REDIS"`
}

var (
	ENV Env
)

func init() {
	envFilePath := "./env/env.json"
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


