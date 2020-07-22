package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Configuration holds config.json info
type Configuration struct {
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration
var logger *log.Logger

// Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a...)
}

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}

	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// version
func version() string {
	return "0.1"
}
