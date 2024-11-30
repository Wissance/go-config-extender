package go_config_extender

import (
	"encoding/json"
	"errors"
	"github.com/ohler55/ojg/jp"
	sf "github.com/wissance/stringFormatter"
	"os"
	"path/filepath"
	"strings"
)

const technicalEnvPrefix = "__"

func LoadJSONConfig[T any](configFile string) (T, error) {
	fileData, err := readJSONConfigStr(configFile)
	var cfg T
	if err = json.Unmarshal(fileData, &cfg); err != nil {
		return nil, errors.New(sf.Format("an error occurred during config file unmarshal:  {0}", err.Error()))
	}
	return cfg, nil
}

func LoadJSONConfigWithEnvOverride[T any](configFile string) (T, error) {
	fileData, err := readJSONConfigStr(configFile)
	if err != nil {
		return nil, err
	}
	var rawCfg map[string]interface{}
	if err = json.Unmarshal(fileData, &rawCfg); err != nil {
		return nil, errors.New(sf.Format("an error occurred during config file unmarshal:  {0}", err.Error()))
	}
	allEnvVars := os.Environ()
	var techEnvVars = map[string]string{}
	for _, pair := range allEnvVars {
		// 1. Pair Key=Value should start from __
		if strings.HasPrefix(pair, technicalEnvPrefix) {
			// 2. Split match pair by =
			parts := strings.Split(pair, "=")
			envVarPath := strings.TrimPrefix(parts[0], technicalEnvPrefix)
			techEnvVars[envVarPath] = parts[1]
		}
	}

	for k, v := range techEnvVars {
		mask, _ := jp.ParseString(k)
		// todo(UMV) -> type ...
		// ???
		// res := mask.Get(rawCfg)
		// understand data type (v), we consider only simple types: bool, int64, float64, datetime, string
		_ = mask.Set(rawCfg, v)
	}

	modifiedData, err := json.Marshal(&rawCfg)
	if err != nil {
		return nil, errors.New(sf.Format("an error occurred during saving applying changes from Env back to JSON : {0}", err.Error()))
	}

	var cfg T
	if err = json.Unmarshal(modifiedData, &cfg); err != nil {
		return nil, errors.New(sf.Format("an error occurred during modified config file unmarshal:  {0}", err.Error()))
	}
	return cfg, nil
}

func readJSONConfigStr(configFile string) ([]byte, error) {
	absPath, err := filepath.Abs(configFile)
	if err != nil {
		return nil, errors.New(sf.Format("an error occurred during getting config file abs path: {0}", err.Error()))
	}

	fileData, err := os.ReadFile(absPath)
	if err != nil {
		return nil, errors.New(sf.Format("an error occurred during config file reading: {0}", err.Error()))
	}

	return fileData, nil
}
