package core

import (
	"os"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

type TorimaConfig struct {
	DefaultOrigin   string   `yaml:"default_origin" default:"127.0.0.1:5000"`
	Host            string   `yaml:"host" default:"http://127.0.0.1:8080"`
	Port            int      `yaml:"port" default:"8080" `
	Scheme          string   `yaml:"scheme" default:"http"`
	SkipAuthList    []string `yaml:"skip_auth_list" default:"[]"`
	ProtectionScope []string `yaml:"protection_scope" default:"[]"`
	WebRoot         string   `yaml:"web_root" default:"/torima"`
}

func ReadConfig() (*TorimaConfig, error) {
	var m TorimaConfig
	var def TorimaConfig // default config

	if err := defaults.Set(&m); err != nil {
		return nil, err
	}

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
