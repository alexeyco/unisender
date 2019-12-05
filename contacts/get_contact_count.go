package contacts

import (
	"fmt"

	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/contacts/getcontactcount/

type GetContactCountRequest interface {
	ParamsTagID(tagID int64) GetContactCountRequest
	ParamsTypeAddress() GetContactCountRequest
	ParamsTypePhone() GetContactCountRequest
	ParamsSearch(search string) GetContactCountRequest
	Execute() (count int64, err error)
}

type GetContactCountResult struct {
	Count int64 `json:"count"`
}

type getContactCountRequest struct {
	request *api.Request
}

func (r *getContactCountRequest) ParamsTagID(tagID int64) GetContactCountRequest {
	r.request.Add("params[tagId]", fmt.Sprintf("%d", tagID))
	return r
}

func (r *getContactCountRequest) ParamsTypeAddress() GetContactCountRequest {
	r.request.Add("params[type]", "address")
	return r
}

func (r *getContactCountRequest) ParamsTypePhone() GetContactCountRequest {
	r.request.Add("params[type]", "phone")
	return r
}

func (r *getContactCountRequest) ParamsSearch(search string) GetContactCountRequest {
	r.request.Add("params[search]", search)
	return r
}

func (r *getContactCountRequest) Execute() (count int64, err error) {
	var res GetContactCountResult
	if err = r.request.Execute("getContactCount", &res); err != nil {
		return
	}

	count = res.Count

	return
}

func GetContactCount(request *api.Request, listID int64) GetContactCountRequest {
	request.Add("list_id", fmt.Sprintf("%d", listID))

	return &getContactCountRequest{
		request: request,
	}
}
