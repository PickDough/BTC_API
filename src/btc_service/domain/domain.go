package domain

type BitcoinRate struct {
	Time     string   `json:"time"`
	Currency Currency `json:"currency"`
}

type Currency struct {
	Code        string  `json:"сode"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}
