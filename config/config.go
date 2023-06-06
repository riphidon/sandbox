package config

import (
	"encoding/json"
	"os"
	"sandbox-api/logs"
)

const configLoaded string = "Successfully Loaded"

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type Config struct {
	Port     int    `json:"port"`
	Env      string `json:"env"`
	Pepper   string `json:"pepper"`
	HMACKey  string `json:"hmac_key"`
	DataBase PostgresConfig
}

func LoadConfig(configReq bool) Config {
	log := logs.NewAppLogger()
	// Open file.
	f, err := os.Open(".config")
	if err != nil {
		// If error opening file, print message
		// notifying we are using default config.
		if configReq {
			log.Startf("[CONFIG] Error : %v", err)
			panic(err)
		}
	}
	var c Config
	dec := json.NewDecoder(f)
	// Object mapped into struct using json tags.
	err = dec.Decode(&c)
	if err != nil {
		log.Startf(" [CONFIG] Error : %v", err)
		panic(err)
	}
	log.Startf(" [CONFIG] %v", configLoaded)
	return c
}
