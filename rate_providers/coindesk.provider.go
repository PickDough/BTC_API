package rate_providers

import (
	"SE_School/models"
	"encoding/json"
	"net/http"
)

type CoindeskRateProvider struct {
}

func (provider *CoindeskRateProvider) GetRate() (*models.BitcoinRate, error) {
	// Getting data from external resource
	resp, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice/UAH.json")
	if err != nil {
		return nil, err
	}

	var btcInfo map[string]interface{}

	if err = json.NewDecoder(resp.Body).Decode(&btcInfo); err != nil {
		return nil, err
	}

	btcRate := &models.BitcoinRate{}

	btcRate.Time = btcInfo["time"].(map[string]interface{})["updated"].(string)
	uah := btcInfo["bpi"].(map[string]interface{})["UAH"].(map[string]interface{})

	btcRate.Currency.Code = uah["code"].(string)
	btcRate.Currency.Rate = uah["rate"].(string)
	btcRate.Currency.Description = uah["description"].(string)
	btcRate.Currency.RateFloat = uah["rate_float"].(float64)

	return btcRate, err
}
