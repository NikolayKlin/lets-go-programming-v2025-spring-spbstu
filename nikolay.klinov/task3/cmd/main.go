package main

import (
	"flag"

	"task3/internal/configuration"
	"task3/internal/currency"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("Config path is required")
	}

	cfg, runErr := configuration.LoadConfig(*configPath)
	if runErr != nil {
		panic(runErr)
	}

	curr, runErr := currency.ParseRatesFromXML(cfg.InputFile)
	if runErr != nil {
		panic(runErr)
	}

	currency.SortByRate(curr)

	runErr = currency.ExportToJSON(curr, cfg.OutputFile)
	if runErr != nil {
		panic(runErr)
	}

}
