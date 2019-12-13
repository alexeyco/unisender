package partners

import (
	"encoding/json"

	"github.com/alexeyco/unisender/api"
)

//
type GetUserInfoResultService struct {
	ID        int64  `json:"id"`
	Tariff    string `json:"tariff"`
	PaidUntil string `json:"paidUntil"`
}

// GetUserInfoResult user info.
type GetUserInfoResult struct {
	Login            string                     `json:"login"`
	Email            string                     `json:"email"`
	FirstName        string                     `json:"firstname"`
	MiddleName       string                     `json:"middlename"`
	LastName         string                     `json:"lastname"`
	RegTime          string                     `json:"regtime"`
	Phone            string                     `json:"phone"`
	Company          string                     `json:"company"`
	Channel          string                     `json:"channel"`
	Timezone         string                     `json:"timezone"`
	Master           string                     `json:"master"`
	Balance          json.Number                `json:"balance"`
	Currency         string                     `json:"currency"`
	EmailsPaid       int64                      `json:"emails_paid"`
	EmailsUsed       int64                      `json:"emails_used"`
	PeriodEmailsPaid int64                      `json:"period_emails_paid"`
	PeriodEmailsUsed int64                      `json:"period_emails_used"`
	EmailPeriodStart api.Time                   `json:"email_period_start"`
	EmailPeriodEnd   api.Time                   `json:"email_period_end"`
	TariffID         json.Number                `json:"tariff_id"`
	NextTariffID     int64                      `json:"next_tariff_id"`
	Services         []GetUserInfoResultService `json:"services"` // TODO clean possible HTML
	Country          string                     `json:"country"`
	Language         string                     `json:"language"`
	Rights           string                     `json:"rights"`
	BalanceBonus     json.Number                `json:"balance_bonus"`
	RegRef           string                     `json:"reg_ref"`
	RegUrl           string                     `json:"reg_url"`
	ApiMode          string                     `json:"api_mode"`
	SubscribersTotal int64                      `json:"subscribers_total"`
	SubscribersUsed  int64                      `json:"subscribers_used"`
}

type GetUserInfoRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetUserInfoRequest) Execute() (res *GetUserInfoResult, err error) {
	var result GetUserInfoResult
	if err = r.request.Execute("getUserInfo", &result); err != nil {
		return
	}

	res = &result

	return
}

// GetUserInfo returns requesst to get user info.
func GetUserInfo(request *api.Request, login string) *GetUserInfoRequest {
	request.Add("login", login)

	return &GetUserInfoRequest{
		request: request,
	}
}
