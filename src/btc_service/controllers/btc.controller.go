package controllers

import (
	"BTC_Service/domain"
	"BTC_Service/utils"
	"net/http"
)

type BtcService interface {
	GetBtcRate() (*domain.BitcoinRate, error)
}

var BtcServ BtcService

func Rate(w http.ResponseWriter, r *http.Request) {
	btcRate, err := BtcServ.GetBtcRate()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.Respond(w, utils.Message("Couldn't fetch bitcoin price"))
		return
	}

	utils.Respond(w, map[string]interface{}{"bitcoin_rate": btcRate})
}
