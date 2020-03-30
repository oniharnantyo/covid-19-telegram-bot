package mathdroid

import (
	"context"
	config2 "covid19-bot/config"
	"fmt"
	"log"
	"testing"
)

func TestGetSummary(t *testing.T) {
	ctx := context.Background()

	config, err := config2.Init("../.config.toml")
	if err != nil {
		log.Fatal(err)
	}

	Init(config.MathdroidURL)

	data, err := GetSummary(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

}

func TestGetConfirmed(t *testing.T) {
	ctx := context.Background()

	config, err := config2.Init("../.config.toml")
	if err != nil {
		log.Fatal(err)
	}

	Init(config.MathdroidURL)

	data, err := GetConfirmed(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

}

func TestGetRecovered(t *testing.T) {
	ctx := context.Background()

	config, err := config2.Init("../.config.toml")
	if err != nil {
		log.Fatal(err)
	}

	Init(config.MathdroidURL)

	data, err := GetRecovered(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

}

func TestGetDeaths(t *testing.T) {
	ctx := context.Background()

	config, err := config2.Init("../.config.toml")
	if err != nil {
		log.Fatal(err)
	}

	Init(config.MathdroidURL)

	data, err := GetDeaths(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

}

func TestGetSummaryPerCountry(t *testing.T) {
	ctx := context.Background()

	config, err := config2.Init("../.config.toml")
	if err != nil {
		log.Fatal(err)
	}

	Init(config.MathdroidURL)

	data, err := GetSummaryPerCountry(ctx, "ID")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

}
