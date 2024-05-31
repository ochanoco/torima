package core

import (
	"fmt"
	"log"
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
	ForceAuthList   []string `yaml:"force_auth_list" default:"[]"`
	ProtectionScope []string `yaml:"protection_scope" default:"[]"`
	WebRoot         string   `yaml:"web_root" default:"/torima"`
}

func readConfig() (*TorimaConfig, error) {
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

func printConfig(config *TorimaConfig) {
	fmt.Println("default_origin:", config.DefaultOrigin)
	fmt.Println("host:", config.Host)
	fmt.Println("port:", config.Port)
	fmt.Println("scheme:", config.Scheme)

	fmt.Println("skip_auth_list:", config.SkipAuthList)

	fmt.Println("protection_scope:", config.ProtectionScope)
	fmt.Println("web_root:", config.WebRoot)
}

func readEnv(name, def string) string {
	value := os.Getenv(name)

	if value == "" {
		fmt.Printf("environment variable '%v' is not found so that proxy use '%v'\n", name, def)
		value = def
	}

	return value
}

func readEnvOrPanic(name string) string {
	value := os.Getenv(name)

	if value == "" {
		log.Fatalf("environment variable '%v' is not found", name)
	}

	return value
}
