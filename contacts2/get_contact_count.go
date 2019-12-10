package contacts2

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// GetContactCountResult request getContactCount request.
type GetContactCountResult struct {
	Count int64 `json:"count"`
}

// GetContactCountRequest request for count contacts by specified conditions.
type GetContactCountRequest struct {
	request *api.Request
}

// ParamsTagID search by tag with a specific id (can be obtained using the getTags method).
func (r *GetContactCountRequest) ParamsTagID(tagID int64) *GetContactCountRequest {
	r.request.Add("params[tagId]", strconv.FormatInt(tagID, 10))
	return r
}

// ParamsTypeAddress search only emails.
func (r *GetContactCountRequest) ParamsTypeAddress(search ...string) *GetContactCountRequest {
	r.request.Add("params[type]", "address")
	if len(search) > 0 {
		r.request.Add("params[search]", search[0])
	}

	return r
}

// ParamsTypePhone search only phones.
func (r *GetContactCountRequest) ParamsTypePhone(search ...string) *GetContactCountRequest {
	r.request.Add("params[type]", "phone")
	if len(search) > 0 {
		r.request.Add("params[search]", search[0])
	}

	return r
}

// Execute sends request to UniSender API and returns result.
func (r *GetContactCountRequest) Execute() (count int64, err error) {
	var res GetContactCountResult
	if err = r.request.Execute("getContactCount", &res); err != nil {
		return
	}

	count = res.Count

	return
}

// GetContactCount returns request for count contacts by specified conditions.
//
// See https://www.unisender.com/en/support/api/contacts/getcontactcount/
func GetContactCount(request *api.Request, listID int64) *GetContactCountRequest {
	request.Add("list_id", strconv.FormatInt(listID, 10))

	return &GetContactCountRequest{
		request: request,
	}
}
