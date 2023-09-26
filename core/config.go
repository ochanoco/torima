package core

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type OchanocoConfig struct {
	DefaultOrigin   string   `yaml:"default_origin"`
	Port            int      `yaml:"port"`
	Scheme          string   `yaml:"scheme"`
	WhiteListPath   []string `yaml:"white_list_path"`
	ProtectionScope []string `yaml:"protection_scope"`
}

func readConfig() (OchanocoConfig, error) {
	f, err := os.Open(CONFIG_FILE)
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()

	d := yaml.NewDecoder(f)

	var m OchanocoConfig
	if err := d.Decode(&m); err != nil {
		log.Fatal(err)
	}

	return m, err
}
