package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type (
	Config struct {
		PublicUrl       string `toml:"public_url"`
		Token           string `toml:"token"`
		Port            string `toml:"port"`
		MathdroidURL    string `toml:"mathdroid_url"`
		LocationiqToken string `toml:"locationiq_token"`
	}
)

func Init(configFile string) (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		fmt.Println(err)
		return Config{}, err
	}

	return config, nil
}
