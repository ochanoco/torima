package core

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type OchanocoConfig struct {
	DefaultOrigin   string   `yaml:"default_origin"`
	Port            int      `yaml:"port"`
	WhiteListPath   []string `yaml:"white_list_path"`
	WhiteListDirs   []string `yaml:"white_list_dirs"`
	AcceptedOrigins []string `yaml:"accepted_origins"`
	IgnoredOrigins  []string `yaml:"ignored_origins"`
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
