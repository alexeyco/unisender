package contacts

import "github.com/alexeyco/unisender/api"

// GetTotalContactsCountResult response of getTotalContactsCount request.
type GetTotalContactsCountResult struct {
	Total int64 `json:"total"`
}

// GetTotalContactsCountRequest request that counts contacts database size by the user login.
type GetTotalContactsCountRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetTotalContactsCountRequest) Execute() (count int64, err error) {
	var result GetTotalContactsCountResult
	if err = r.request.Execute("getTotalContactsCount", &result); err != nil {
		return
	}

	count = result.Total

	return
}

// GetTotalContactsCount returns request that counts contacts database size by the user login.
func GetTotalContactsCount(request *api.Request, login string) *GetTotalContactsCountRequest {
	request.Add("login", login)

	return &GetTotalContactsCountRequest{
		request: request,
	}
}
