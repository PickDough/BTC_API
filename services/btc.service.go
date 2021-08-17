package services

import (
	"SE_School/models"
)

type BtcRateProvider interface {
	GetRate() (*models.BitcoinRate, error)
}

type BtcService struct {
	RateProvider BtcRateProvider
}

func (service *BtcService) GetBtcRate() (*models.BitcoinRate, error) {
	return service.RateProvider.GetRate()
}
