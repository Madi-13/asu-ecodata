package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type (
	Config struct {
		Keycloak    Keycloak `json:"keycloak"`
		Http        HTTP     `json:"http"`
		DbURL       string   `json:"db_url"`
		Environment string   `json:"environment"`
	}
	HTTP struct {
		Host               string        `json:"host"`
		Port               string        `json:"port"`
		ReadTimeout        time.Duration `json:"read_timeout"`
		WriteTimeout       time.Duration `json:"write_timeout"`
		MaxHeaderMegabytes int           `json:"max_header_megabytes"`
	}
	Keycloak struct {
		Host      string `json:"host"`
		Realm     string `json:"realm"`
		Client    string `json:"client"`
		Secret    string `json:"secret"`
		PublicKey string `json:"public_key"`
	}
)

func LoadLocalConfig() (*Config, error) {
	var Conf = &Config{}
	var path string
	projectDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	flag.StringVar(&path, "config", fmt.Sprintf("%s/config/conf.json", projectDir), "Config path")
	flag.Parse()
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Conf)
	if err != nil {
		return nil, err
	}
	return Conf, nil
}
