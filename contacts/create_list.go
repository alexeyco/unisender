package contacts

import (
	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/partners/createlist/

type CreateListRequest interface {
	BeforeSubscribeUrl(u string) CreateListRequest
	AfterSubscribeUrl(u string) CreateListRequest
	Execute() (int64, error)
}

type CreateListResponse struct {
	ID int64 `json:"id"`
}

type createListRequest struct {
	request *api.Request
}

func (r *createListRequest) BeforeSubscribeUrl(u string) CreateListRequest {
	r.request.Add("before_subscribe_url", u)
	return r
}

func (r *createListRequest) AfterSubscribeUrl(u string) CreateListRequest {
	r.request.Add("after_subscribe_url", u)
	return r
}

func (r *createListRequest) Execute() (listID int64, err error) {
	var res CreateListResponse
	if err = r.request.Execute("createList", &res); err != nil {
		return
	}

	listID = res.ID

	return
}

func CreateList(request *api.Request, title string) CreateListRequest {
	request.Add("title", title)

	return &createListRequest{
		request: request,
	}
}
