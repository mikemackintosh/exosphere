package main

import (
	"exo"
	"flag"
	"log"
)

var (
	flagConfig string
)

func init() {
	flag.StringVar(&flagConfig, "c", "example.hcl", "Coonfig")
}

func main() {
	var cfg exo.Config
	if err := exo.LoadConfig("example.hcl", &cfg); err != nil {
		log.Fatal(err)
	}

	if err := exo.StartServers(&cfg); err != nil {
		log.Fatal(err)
	}
}
