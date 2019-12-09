package contacts

import "github.com/alexeyco/unisender/api"

// GetTotalContactsCountResponse response of getTotalContactsCount request.
type GetTotalContactsCountResponse struct {
	Total int64 `json:"total"`
}

// GetTotalContactsCountRequest request that counts contacts database size by the user login.
type GetTotalContactsCountRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetTotalContactsCountRequest) Execute() (count int64, err error) {
	var res GetTotalContactsCountResponse
	if err = r.request.Execute("getTotalContactsCount", &res); err != nil {
		return
	}

	count = res.Total

	return
}

// GetTotalContactsCount returns request that counts contacts database size by the user login.
//
// See https://www.unisender.com/en/support/api/partners/gettotalcontactscount/
func GetTotalContactsCount(request *api.Request, login string) *GetTotalContactsCountRequest {
	request.Add("login", login)

	return &GetTotalContactsCountRequest{
		request: request,
	}
}
