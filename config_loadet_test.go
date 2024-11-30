package go_config_extender_test

import (
	"github.com/stretchr/testify/assert"
	gce "github.com/wissance/go-config-extender"
	sf "github.com/wissance/stringFormatter"
	"os"
	"path"
	"testing"
)

type testServerConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type testHttpLoggingConfig struct {
	Enabled bool `json:"enabled"`
}

type testLoggingConfig struct {
	Level   string                `json:"level"`
	HttpLog testHttpLoggingConfig `json:"httpLog"`
}

type testSensorsConfig struct {
	Threshold float64 `json:"threshold"`
}

type testConfig struct {
	Server  testServerConfig  `json:"server"`
	Logging testLoggingConfig `json:"logging"`
	Sensors testSensorsConfig `json:"sensors"`
}

func TestLoadJSONConfigWithEnvOverride(t *testing.T) {
	// 1. Set some tech vars
	portEnvVar := "__server.port"
	httpLoggingEnabledEnvVar := "__logging.httpLog.enabled"
	sensorsThresholdEnvVar := "__sensors.threshold"
	expectedPortValue := 6000
	expectedEnabledValues := true
	expectedThresholdValue := 0.25
	err := addTechEnvVarForTest(portEnvVar, sf.Format("{0}", expectedPortValue))
	assert.NoError(t, err)
	err = addTechEnvVarForTest(httpLoggingEnabledEnvVar, sf.Format("{0}", expectedEnabledValues))
	assert.NoError(t, err)
	err = addTechEnvVarForTest(sensorsThresholdEnvVar, sf.Format("{0}", expectedThresholdValue))
	assert.NoError(t, err)
	// 2. Load config, check envs were applied
	configFile := path.Join(".", "testConfigs", "testConfig1.json")
	cfg, err := gce.LoadJSONConfigWithEnvOverride[testConfig](configFile)
	assert.NoError(t, err)
	assert.Equal(t, expectedPortValue, cfg.Server.Port)
	assert.Equal(t, expectedEnabledValues, cfg.Logging.HttpLog.Enabled)
	assert.Equal(t, expectedThresholdValue, cfg.Sensors.Threshold)
	// 3. Drop all tech vars
	_ = removeTechEnvVarForTest(portEnvVar)
	_ = removeTechEnvVarForTest(httpLoggingEnabledEnvVar)
	_ = removeTechEnvVarForTest(sensorsThresholdEnvVar)
}

func addTechEnvVarForTest(key string, value string) error {
	return os.Setenv(key, value)
}

func removeTechEnvVarForTest(key string) error {
	return os.Unsetenv(key)
}
