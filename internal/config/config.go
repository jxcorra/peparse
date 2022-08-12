package config

import (
	"encoding/json"
	"fmt"

	common "github.com/jxcorra/peparse/internal/common"
)

func ParseConfiguration(configData []byte) (common.Resources, error) {
	var resources common.Resources
	err := json.Unmarshal(configData, &resources)
	if err != nil {
		return common.Resources{}, fmt.Errorf("configuration content cannot be serialized")
	}

	return resources, nil
}
