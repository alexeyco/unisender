package contacts

import (
	"fmt"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/contacts/subscribe/

type SubscribeRequest interface {
	Email(email string) SubscribeRequest
	Phone(phone string) SubscribeRequest
	Name(name string) SubscribeRequest
	Tags(tags ...string) SubscribeRequest
	DoubleOptin() SubscribeRequest
	Overwrite() SubscribeRequest
	Execute() (personID int64, err error)
}

type subscribeResponse struct {
	PersonID int64 `json:"person_id"`
}

type subscribeRequest struct {
	request *api.Request
}

func (r *subscribeRequest) Email(email string) SubscribeRequest {
	r.request.Add("fields[email]", email)
	return r
}

func (r *subscribeRequest) Phone(phone string) SubscribeRequest {
	r.request.Add("fields[phone]", phone)
	return r
}

func (r *subscribeRequest) Name(name string) SubscribeRequest {
	r.request.Add("fields[Name]", name)
	return r
}

func (r *subscribeRequest) Tags(tags ...string) SubscribeRequest {
	r.request.Add("tags", strings.Join(tags, ","))
	return r
}

func (r *subscribeRequest) DoubleOptin() SubscribeRequest {
	r.request.Add("double_optin", "0")
	return r
}

func (r *subscribeRequest) Overwrite() SubscribeRequest {
	r.request.Add("overwrite", "0")
	return r
}

func (r *subscribeRequest) Execute() (personID int64, err error) {
	var res subscribeResponse
	if err = r.request.Execute("subscribe", &res); err != nil {
		return
	}

	personID = res.PersonID

	return
}

func Subscribe(request *api.Request, listIDs ...int64) SubscribeRequest {
	ids := make([]string, len(listIDs))
	for n, id := range listIDs {
		ids[n] = fmt.Sprintf("%d", id)
	}

	request.Add("list_ids", strings.Join(ids, ","))

	return &subscribeRequest{
		request: request,
	}
}
