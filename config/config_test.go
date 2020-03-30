package config

import (
	"fmt"
	"log"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config, err := Init("../.config.toml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
}
