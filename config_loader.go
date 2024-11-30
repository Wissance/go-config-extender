package go_config_extender

import (
	"encoding/json"
	"errors"
	sf "github.com/wissance/stringFormatter"
	"os"
	"path/filepath"
)

func LoadJSONConfig[T any](configFile string) (T, error) {
	absPath, err := filepath.Abs(configFile)
	if err != nil {
		return nil, errors.New(sf.Format("an error occurred during getting config file abs path: {0}", err.Error()))
	}

	fileData, err := os.ReadFile(absPath)
	if err != nil {
		return nil, errors.New(sf.Format("an error occurred during config file reading: {0}", err.Error()))
	}
	var cfg T
	if err = json.Unmarshal(fileData, &cfg); err != nil {
		return nil, errors.New(sf.Format("an error occurred during config file unmarshal:  {0}", err.Error()))
	}
	return cfg, nil
}

func LoadJSONConfigWithEnvOverride[T any](file string) (T, error) {
	return nil, nil
}
