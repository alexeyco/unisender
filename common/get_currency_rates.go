package common

import "github.com/alexeyco/unisender/api"

// CurrencyRate currency rate.
type CurrencyRate struct {
	ID            int64   `json:"id"`
	Code          string  `json:"code"`
	RateToUSD     float64 `json:"rate_to_usd,string"`
	MinPaymentSum float64 `json:"min_payment_sum,string"`
	IsVisible     bool    `json:"is_visible"`
}

// GetCurrencyRatesRequest request to get currency rate.
type GetCurrencyRatesRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetCurrencyRatesRequest) Execute() (res []CurrencyRate, err error) {
	err = r.request.Execute("getCurrencyRates", &res)
	return
}

// GetCurrencyRates allows you to get a list of all currencies in the UniSender system.
func GetCurrencyRates(request *api.Request) *GetCurrencyRatesRequest {
	request.Add("format", "json")

	return &GetCurrencyRatesRequest{
		request: request,
	}
}
