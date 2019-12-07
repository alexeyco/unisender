package contacts

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/contacts/subscribe/

type SubscribeRequest interface {
	Field(field, value string) SubscribeRequest
	Email(email string) SubscribeRequest
	Phone(phone string) SubscribeRequest
	Tags(tags ...string) SubscribeRequest
	DoubleOptinUnconfirmed() SubscribeRequest
	DoubleOptinConfirmed() SubscribeRequest
	DoubleOptinConfirmedIfActiveOrNew() SubscribeRequest
	DoNotOverwrite() SubscribeRequest
	OverwriteAll() SubscribeRequest
	OverwritePartially() SubscribeRequest
	Execute() (personID int64, err error)
}

type SubscribeResponse struct {
	PersonID int64 `json:"person_id"`
}

type subscribeRequest struct {
	request *api.Request
}

func (r *subscribeRequest) Field(field, value string) SubscribeRequest {
	r.request.Add(fmt.Sprintf("fields[%s]", field), value)
	return r
}

func (r *subscribeRequest) Email(email string) SubscribeRequest {
	return r.Field("email", email)
}

func (r *subscribeRequest) Phone(phone string) SubscribeRequest {
	return r.Field("phone", phone)
}

func (r *subscribeRequest) Tags(tags ...string) SubscribeRequest {
	r.request.Add("tags", strings.Join(tags, ","))
	return r
}

func (r *subscribeRequest) DoubleOptinUnconfirmed() SubscribeRequest {
	r.request.Add("double_optin", "0")
	return r
}

func (r *subscribeRequest) DoubleOptinConfirmed() SubscribeRequest {
	r.request.Add("double_optin", "3")
	return r
}

func (r *subscribeRequest) DoubleOptinConfirmedIfActiveOrNew() SubscribeRequest {
	r.request.Add("double_optin", "4")
	return r
}

func (r *subscribeRequest) DoNotOverwrite() SubscribeRequest {
	r.request.Add("overwrite", "0")
	return r
}

func (r *subscribeRequest) OverwriteAll() SubscribeRequest {
	r.request.Add("overwrite", "1")
	return r
}

func (r *subscribeRequest) OverwritePartially() SubscribeRequest {
	r.request.Add("overwrite", "2")
	return r
}

func (r *subscribeRequest) Execute() (personID int64, err error) {
	var res SubscribeResponse
	if err = r.request.Execute("subscribe", &res); err != nil {
		return
	}

	personID = res.PersonID

	return
}

func Subscribe(request *api.Request, listIDs ...int64) SubscribeRequest {
	ids := make([]string, len(listIDs))
	for n, id := range listIDs {
		ids[n] = strconv.FormatInt(id, 10)
	}

	request.Add("list_ids", strings.Join(ids, ","))

	return &subscribeRequest{
		request: request,
	}
}
