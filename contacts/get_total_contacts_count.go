package contacts

import "github.com/alexeyco/unisender/api"

// See https://www.unisender.com/en/support/api/partners/gettotalcontactscount/

type GetTotalContactsCountRequest interface {
	Execute() (count int64, err error)
}

type GetTotalContactsCountResponse struct {
	Total int64 `json:"total"`
}

type getTotalContactsCountRequest struct {
	request *api.Request
}

func (r *getTotalContactsCountRequest) Execute() (count int64, err error) {
	var res GetTotalContactsCountResponse
	if err = r.request.Execute("getTotalContactsCount", &res); err != nil {
		return
	}

	count = res.Total

	return
}

func GetTotalContactsCount(request *api.Request, login string) GetTotalContactsCountRequest {
	request.Add("login", login)

	return &getTotalContactsCountRequest{
		request: request,
	}
}
