package config

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BaseDir string `yaml:"baseDir"`
}

func Configure(args string) Config {
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "inbx", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {

		if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
			log.Fatal(err)
		}

		if f, err := os.Create(configPath); err != nil {
			log.Fatal(err)
		} else {
			defer f.Close()
			f.WriteString("baseDir: ~/inbx\n")
		}

	}

	f, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("failed to open config file: %s", err)
	}

	defer f.Close()

	var config Config
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil && err != io.EOF {
		log.Fatalf("failed to parse config file: %s", err)
	}

	return config
}
