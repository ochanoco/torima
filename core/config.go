package core

import (
	"fmt"
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

type OchanocoConfig struct {
	DefaultOrigin   string   `yaml:"default_origin" default:"127.0.0.1:8080"`
	Port            int      `yaml:"port" default:"8080" `
	Scheme          string   `yaml:"scheme" default:"https"`
	WhiteListPath   []string `yaml:"white_list_path" default:"[]"`
	ProtectionScope []string `yaml:"protection_scope" default:"[]"`
}

func readConfig() (*OchanocoConfig, error) {
	var m OchanocoConfig
	var def OchanocoConfig // default config

	if err := defaults.Set(&def); err != nil {
		return nil, err
	}

	f, err := os.Open(CONFIG_FILE)
	if err != nil {
		return &def, err
	}
	defer f.Close()

	d := yaml.NewDecoder(f)
	if err := d.Decode(&m); err != nil {
		return &def, err
	}

	return &m, err
}

func printConfig(config *OchanocoConfig) {
	fmt.Println("default_origin:", config.DefaultOrigin)
	fmt.Println("port:", config.Port)
	fmt.Println("scheme:", config.Scheme)
	fmt.Println("white_list_path:", config.WhiteListPath)
	fmt.Println("protection_scope:", config.ProtectionScope)
}

func readEnv(name, def string) string {
	value := os.Getenv(name)

	if value == "" {
		fmt.Printf("environment variable '%v' is not found so that proxy use '%v'\n", name, def)
		value = def
	}

	return value
}
