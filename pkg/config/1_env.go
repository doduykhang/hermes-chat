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
	Redis Redis `json:"REDIS"`
}

var (
	ENV Env
)

func init() {
	bytes, err := os.ReadFile("./env/rc.json")
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


