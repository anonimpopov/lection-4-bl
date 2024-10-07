package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Port  string `json:"port"`
	Debug bool   `json:"debug"`
}

func main() {
	var cfg Config
	var cfgPath string

	cfgPath = filepath.Clean(cfgPath)
	file, err := os.Open(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", cfg)
}
