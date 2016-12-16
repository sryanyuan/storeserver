package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// AppConfig is config struct for app
type AppConfig struct {
	HTTPAddress string   `validate:"nonzero"`
	StoreSource string `validate:"nonzero"`
	AuthUser    []string `validate:"nonzero"`
}

func loadConfigFromFile(file string) (*AppConfig, error) {
	f, err := os.Open(file)
	if nil != err {
		log.Println("[ERR] Can't open config file:", file)
		return nil, err
	}
	defer f.Close()

	fileBytes, err := ioutil.ReadAll(f)
	if nil != err {
		log.Println("[ERR] Can't read config file:", file)
		return nil, err
	}

	//	parse the json
	var config AppConfig
	if err = json.Unmarshal(fileBytes, &config); nil != err {
		return nil, err
	}

	return &config, err
}
