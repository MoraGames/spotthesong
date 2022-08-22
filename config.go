package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const (
	applicationConfig Path = "config.yaml"
)

type Config struct {
	SPOTIFY_ID string
	SPOTIFY_SECRET string
}

func (c *Config) GetConfig(){
	yamlFile, err := ioutil.ReadFile(applicationConfig.ToString())
	if err != nil {
		log.Panicf("getConfig: yaml readfile error: %v\n", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Panicf("getConfig: yaml unmarshal error: %v\n", err)
	}
}