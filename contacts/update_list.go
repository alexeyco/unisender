package contacts

import (
	"fmt"

	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/partners/updatelist/

type UpdateListRequest interface {
	BeforeSubscribeUrl(u string) UpdateListRequest
	AfterSubscribeUrl(u string) UpdateListRequest
	Execute() error
}

type updateListRequest struct {
	request *api.Request
}

func (r *updateListRequest) BeforeSubscribeUrl(u string) UpdateListRequest {
	r.request.Add("before_subscribe_url", u)
	return r
}

func (r *updateListRequest) AfterSubscribeUrl(u string) UpdateListRequest {
	r.request.Add("after_subscribe_url", u)
	return r
}

func (r *updateListRequest) Execute() error {
	return r.request.Execute("updateList", nil)
}

func UpdateList(request *api.Request, listID int64, title string) UpdateListRequest {
	request.Add("list_id", fmt.Sprintf("%d", listID)).
		Add("title", title)

	return &updateListRequest{
		request: request,
	}
}
