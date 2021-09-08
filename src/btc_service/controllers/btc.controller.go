package controllers

import (
	"BTC_Service/domain"
	"BTC_Service/utils"
	"fmt"
	"net/http"
)

type BtcService interface {
	GetBtcRate() (*domain.BitcoinRate, error)
}

type MessageBroker interface {
	SendInfo(msg string)
	SendWarning(msg string)
	SendError(msg string)
	Close()
}

var BtcServ BtcService
var MsgBroker MessageBroker

func Rate(w http.ResponseWriter, r *http.Request) {
	MsgBroker.SendInfo(fmt.Sprintf("%s", "Request at \\btcRate"))

	btcRate, err := BtcServ.GetBtcRate()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		MsgBroker.SendError(fmt.Sprintf("btc_service: %s", "Couldn't fetch bitcoin price"))
		utils.Respond(w, utils.Message("Couldn't fetch bitcoin price"))
		return
	}

	utils.Respond(w, map[string]interface{}{"bitcoin_rate": btcRate})
}
