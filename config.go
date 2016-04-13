// iotback project iotback.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const configFileName = "config.json"

type aConfig struct {
	username string
	password string
	machine  string
}

func (c aConfig) getaConfig() aConfig {
	f, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Println("config.json not found, creating...")
		t, _ := json.Marshal(c)
		ioutil.WriteFile(configFileName, t, 0777)
	}
	// Need to add error handling here
	json.Unmarshal(f, c)
	return c
}
