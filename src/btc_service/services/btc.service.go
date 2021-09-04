package services

import (
	"BTC_Service/domain"
)

type BtcRateProvider interface {
	GetRate() (*domain.BitcoinRate, error)
}

type BtcService struct {
	RateProvider BtcRateProvider
}

func (service *BtcService) GetBtcRate() (*domain.BitcoinRate, error) {
	return service.RateProvider.GetRate()
}
