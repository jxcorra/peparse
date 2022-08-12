package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	common "github.com/jxcorra/peparse/internal/common"
	config "github.com/jxcorra/peparse/internal/config"
	worker "github.com/jxcorra/peparse/internal/worker"
)

func main() {
	var period int
	flag.IntVar(&period, "period", 5000, "How often run parsing tasks in miliseconds")

	var numOfWorkers int
	flag.IntVar(&numOfWorkers, "workers", 10, "Number of parallel parse tasks")

	var configuration string
	flag.StringVar(&configuration, "configuration", "configuration.json", "Path to configuration json")

	flag.Parse()

	tasks := make(chan common.ResourceConfig, numOfWorkers)
	output := make(chan common.Parsed, numOfWorkers)

	file, err := os.Open(configuration)
	if err != nil {
		panic(fmt.Errorf("no such file %s", configuration))
	}
	defer file.Close()

	configData, err := ioutil.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("data from file %s cannot be read", configuration))
	}

	resources, err := config.ParseConfiguration(configData)
	if err != nil {
		panic(err.Error())
	}

	parameters := common.Parameters{
		Period:       period,
		NumOfWorkers: numOfWorkers,
		Resources:    &resources,
		Tasks:        tasks,
		Output:       output,
	}

	worker.RunPeriodicTask(parameters)
}
