package partners

import (
	"strconv"
	"time"

	"github.com/alexeyco/unisender/api"
)

// ChangeTariffResult request to change tariff result.
type ChangeTariffResult struct {
	TariffID     int64     `json:"tariff_id"`
	ChargedSum   float64   `json:"charged_sum"`
	ChargedBonus float64   `json:"charged_bonus"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
}

// ChangeTariffRequest request to change tariff.
type ChangeTariffRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *ChangeTariffRequest) Execute() (res *ChangeTariffResult, err error) {
	var result ChangeTariffResult
	if err = r.request.Execute("changeTariff", &result); err != nil {
		return
	}

	res = &result

	return
}

// ChangeTariff returns request to change tariff.
func ChangeTariff(request *api.Request, login string, tariffID int64) *ChangeTariffRequest {
	request.Add("login", login).
		Add("tariff_id", strconv.FormatInt(tariffID, 10))

	return &ChangeTariffRequest{
		request: request,
	}
}
