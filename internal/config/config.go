package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	common "github.com/jxcorra/peparse/internal/common"
)

func ParseConfiguration(configPath string) (*common.Resources, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("no such file %s", configPath)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("data from file %s cannot be read", configPath)
	}

	var resources common.Resources
	err = json.Unmarshal(data, &resources)
	if err != nil {
		return nil, fmt.Errorf("content from %s cannot be serialized", configPath)
	}

	return &resources, nil
}
