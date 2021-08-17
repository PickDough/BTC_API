package rate_providers

import "testing"

func TestCoindeskRateProvider_GetRateWilReturnNotEmptyRate(t *testing.T) {
	provider := CoindeskRateProvider{}

	btcRate, err := provider.GetRate()
	if err != nil {
		t.Errorf(err.Error())
	}

	if btcRate == nil || btcRate.Currency.Rate == "" {
		t.Errorf("didn't receive bitcoin rate")
	}
}
