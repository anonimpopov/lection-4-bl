package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var configPath string
	args := os.Args

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.StringVar(&configPath, "c", "config.yaml", "set path to config")

	err := flags.Parse(args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
