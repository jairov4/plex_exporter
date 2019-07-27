package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Token   string `json:"token"`
}

func Load(file string) (*Config, error) {

	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("Config file %s does not exist", file)
		}
	}

	data, _ := ioutil.ReadFile(file)

	var config Config
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		return nil, fmt.Errorf("Config file %s is not a valid json", file)
	}
	return &config, nil
}
