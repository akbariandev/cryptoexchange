package config

import (
	"flag"
	"fmt"
	"gitlab.com/hotelian-company/challenge/internal/entity"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Service    *ServiceConfig             `yaml:"core"`
	Currencies map[string]entity.Currency `yaml:"currencies"`
	Providers  map[string]entity.Provider `yaml:"providers"`
}

type ServiceConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func GetConfig() (*Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func getConfigPath() (string, error) {
	var configPath string
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}
