package core

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TEST_CONFIG = `
port: 9000

skip_auth_list:
  - /favicon.ico

default_origin: 127.0.0.1:9000

protection_scope:
  - example.com

scheme: http
`

func readTestConfig(t *testing.T) (*TorimaConfig, *os.File, error) {
	file, err := os.CreateTemp("", "config.yaml")
	assert.NoError(t, err)

	CONFIG_FILE = file.Name()

	_, err = file.Write([]byte(TEST_CONFIG))
	assert.NoError(t, err)

	config, err := readConfig()

	return config, file, err
}

// test for readConfig
func TestReadConfig(t *testing.T) {
	config, file, err := readTestConfig(t)
	defer os.Remove(file.Name())

	assert.NoError(t, err)

	assert.Equal(t, 9000, config.Port)
	assert.Equal(t, "127.0.0.1:9000", config.DefaultOrigin)
	assert.Equal(t, "http", config.Scheme)
	assert.Equal(t, "/favicon.ico", config.SkipAuthList[0])
	assert.Equal(t, "example.com", config.ProtectionScope[0])
}

func TestReadConfigDefault(t *testing.T) {
	CONFIG_FILE = ""

	config, err := readConfig()

	assert.Error(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "127.0.0.1:5000", config.DefaultOrigin)
	assert.Equal(t, "http://127.0.0.1:8080", config.Host)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, "http", config.Scheme)
	assert.Equal(t, 0, len(config.SkipAuthList))
	assert.Equal(t, 0, len(config.ProtectionScope))
	assert.Equal(t, "/torima", config.WebRoot)
}

// test for readEnv
func TestReadEnv(t *testing.T) {
	os.Setenv("TORIMA_TEST1", "TEST")

	env := readEnv("TORIMA_TEST1", "TEST")
	assert.Equal(t, "TEST", env)

	env = readEnv("TORIMA_TEST2", "TEST")
	assert.Equal(t, "TEST", env)
}
