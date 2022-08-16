package config

import (
	"encoding/json"
	"fmt"

	common "github.com/jxcorra/peparse/internal/common"
)

func checkMandatoryFields(resources *common.Resources) error {
	for _, resource := range resources.Resources {
		if resource.Url == "" {
			return fmt.Errorf("no resource url")
		}

		for _, search := range resource.Search {
			if search.Key.Element == "" {
				return fmt.Errorf("no search key element")
			}
		}
	}

	return nil
}

func ParseConfiguration(configData []byte) (common.Resources, error) {
	var resources common.Resources
	err := json.Unmarshal(configData, &resources)
	if err != nil {
		return common.Resources{}, fmt.Errorf("configuration content cannot be serialized")
	}
	if len(resources.Resources) == 0 {
		return common.Resources{}, fmt.Errorf("no resources found in configuration")
	}

	err = checkMandatoryFields(&resources)
	if err != nil {
		return common.Resources{}, err
	}

	return resources, nil
}

func NewCommunication(numOfWorkers int) common.Communication {
	done := make(chan bool, 1)
	printerDone := make(chan bool, 1)
	workersDone := make(chan bool, numOfWorkers)

	return common.Communication{
		Done: done,
		PrinterDone: printerDone,
		WorkerDone: workersDone,
	}
}
