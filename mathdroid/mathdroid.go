package mathdroid

import (
	"context"
	"covid19-bot/models"
	"errors"
)

func GetSummary(ctx context.Context) (models.SummaryModel, error) {
	extraUrl := ""

	var base models.SummaryModel
	err := Get(ctx, extraUrl, nil, &base)
	if err != nil {
		return models.SummaryModel{}, err
	}

	return base, nil
}

func GetConfirmed(ctx context.Context) ([]models.Detail, error) {
	extraUrl := "/confirmed"

	var details []models.Detail
	err := Get(ctx, extraUrl, nil, &details)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func GetRecovered(ctx context.Context) ([]models.Detail, error) {
	extraUrl := "/recovered"

	var details []models.Detail
	err := Get(ctx, extraUrl, nil, &details)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func GetDeaths(ctx context.Context) ([]models.Detail, error) {
	extraUrl := "/deaths"

	var details []models.Detail
	err := Get(ctx, extraUrl, nil, &details)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func GetSummaryPerCountry(ctx context.Context, country string) (models.SummaryPerCountry, error) {
	if country == "" {
		return models.SummaryPerCountry{}, errors.New("Invalid country parameter")
	}

	extraUrl := "/countries/" + country

	var summaryPerCountry models.SummaryPerCountry
	err := Get(ctx, extraUrl, nil, &summaryPerCountry)
	if err != nil {
		return models.SummaryPerCountry{}, err
	}

	return summaryPerCountry, nil
}
