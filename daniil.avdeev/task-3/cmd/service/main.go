package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
  "sort"

	"github.com/realFrogboy/task-3/internal/configparser"
	"github.com/realFrogboy/task-3/internal/cursesparser"
	"github.com/realFrogboy/task-3/internal/jsonemitter"
)

func main() {
	configFilePath := flag.String("config", "configs/config1.yml", "Path to the config file")
	flag.Parse()

	configFile, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't read config file: %s", err)
		return
	}

	config, err := configparser.Parse(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse config file: %s", err)
		return
	}

	cursesFile, err := os.Open(config.InputFile)
	if err != nil {
 		fmt.Fprintf(os.Stderr, "Can't open input file: %s", err)
		return
	}
	defer cursesFile.Close()

	valutes, err := cursesparser.Parse(cursesFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant parse input file: %s", err)
		return
	}

  sort.Slice(valutes, func(i, j int) bool { return valutes[i].Value > valutes[j].Value })

	outputFile, err := os.OpenFile(config.OutputFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't open output file: %s", err)
		return
	}
	defer outputFile.Close()

	err = jsonemitter.Emit(outputFile, valutes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't emit valutes: %s", err)
		return
	}
}
