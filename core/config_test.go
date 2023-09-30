package core

import (
	"io/ioutil"
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

func readTestConfig(t *testing.T) (OchanocoConfig, error) {
	file, err := ioutil.TempFile("dir", "tmp.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	file.Write([]byte(TEST_CONFIG))

	CONFIG_FILE = "../test/ochanoco.yml"

	config, err := readConfig()

	return config, err
}

// test for readConfig
func TestReadConfig(t *testing.T) {
	config, err := readTestConfig(t)
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
	file, err := ioutil.TempFile("dir", "tmp.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	config, err := readConfig()
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

	if len(config.WhiteListPath) == 0 {
		t.Fatalf("readConfig() is failed (config.WhiteListPath)")
	}

	if len(config.ProtectionScope) == 0 {
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
