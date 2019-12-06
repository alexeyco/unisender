package contacts

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/partners/unsubscribe/

type UnsubscribeRequest interface {
	ContactTypeEmail() UnsubscribeRequest
	ContactTypePhone() UnsubscribeRequest
	ListIDs(listID ...int64) UnsubscribeRequest
	Execute() (err error)
}

type unsubscribeRequest struct {
	request *api.Request
}

func (r *unsubscribeRequest) ContactTypeEmail() UnsubscribeRequest {
	r.request.Add("contact_type", "email")
	return r
}

func (r *unsubscribeRequest) ContactTypePhone() UnsubscribeRequest {
	r.request.Add("contact_type", "phone")
	return r
}

func (r *unsubscribeRequest) ListIDs(listID ...int64) UnsubscribeRequest {
	ids := make([]string, len(listID))
	for n, id := range listID {
		ids[n] = strconv.FormatInt(id, 10)
	}

	r.request.Add("list_ids", strings.Join(ids, ","))

	return r
}

func (r *unsubscribeRequest) Execute() (err error) {
	return r.request.Execute("unsubscribe", nil)
}

func Unsubscribe(request *api.Request, contact string) UnsubscribeRequest {
	request.Add("contact", contact)

	return &unsubscribeRequest{
		request: request,
	}
}
