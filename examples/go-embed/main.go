package main

import (
	"exo"
	"log"

	_ "embed"
)

var (
	//go:embed example.hcl
	config []byte
)

func main() {
	var cfg exo.Config
	if err := exo.LoadConfigFromBytes(config, &cfg); err != nil {
		log.Fatal(err)
	}

	if err := exo.StartServers(&cfg); err != nil {
		log.Fatal(err)
	}
}
