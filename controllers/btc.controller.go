package controllers

import (
	"SE_School/models"
	"SE_School/utils"
	"net/http"
)

type BtcService interface {
	GetBtcRate() (*models.BitcoinRate, error)
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
