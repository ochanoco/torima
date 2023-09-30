package core

import (
	"fmt"
	"os"
	"testing"
)

var TEST_CONFIG = `
port: 9000

white_list_path:
  - /favicon.ico

default_origin: 127.0.0.1:9000

protection_scope:
  - example.com

scheme: http
`

func readTestConfig(t *testing.T) (OchanocoConfig, *os.File, error) {
	file, err := os.CreateTemp("", "config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	CONFIG_FILE = file.Name()

	_, err = file.Write([]byte(TEST_CONFIG))
	if err != nil {
		t.Fatalf("failed to read config (%v) %v", CONFIG_FILE, err)
	}

	config, err := readConfig()

	return config, file, err
}

// test for readConfig
func TestReadConfig(t *testing.T) {
	config, file, err := readTestConfig(t)
	defer os.Remove(file.Name())

	if err != nil {
		t.Fatalf("readConfig() is failed: %v", err)
	}

	if config.Port != 9000 {
		t.Fatalf("readConfig() is failed")
	}

	if config.DefaultOrigin != "127.0.0.1:9000" {
		t.Fatalf("readConfig() is failed (config.DefaultOrigin)")
	}

	if config.Scheme != "http" {
		t.Fatalf("readConfig() is failed (config.Scheme)")
	}

	if config.WhiteListPath[0] != "/favicon.ico" {
		t.Fatalf("readConfig() is failed (config.WhiteListPath)")
	}

	if config.ProtectionScope[0] != "example.com" {
		t.Fatalf("readConfig() is failed (config.ProtectionScope)")
	}
}

func TestReadConfigDefault(t *testing.T) {
	CONFIG_FILE = ""

	config, err := readConfig()
	fmt.Printf("%v %v", config, len(config.WhiteListPath))
	if err != nil {
		t.Fatalf("readConfig() is failed: %v", err)
	}

	if config.DefaultOrigin != "127.0.0.1:8080" {
		t.Fatalf("readConfig() is failed (config.DefaultOrigin)")
	}

	if config.Port != 8080 {
		t.Fatalf("readConfig() is failed (config.Port)")
	}

	if config.Scheme != "https" {
		t.Fatalf("readConfig() is failed (config.Scheme)")
	}

	if len(config.WhiteListPath) != 0 {
		t.Fatalf("readConfig() is failed (config.WhiteListPath)")
	}

	if len(config.ProtectionScope) != 0 {
		t.Fatalf("readConfig() is failed (config.ProtectionScope)")
	}
}

// test for readEnv
func TestReadEnv(t *testing.T) {
	os.Setenv("OCHANOCO_TEST1", "TEST")

	env := readEnv("OCHANOCO_TEST1", "TEST")

	if env != "TEST" {
		t.Fatalf("readEnv() is failed")
	}

	env = readEnv("OCHANOCO_TEST2", "TEST")

	if env != "TEST" {
		t.Fatalf("readEnv() is failed")
	}
}
