package contacts

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/partners/exclude/

type ExcludeRequest interface {
	ContactTypeEmail() ExcludeRequest
	ContactTypePhone() ExcludeRequest
	ListIDs(listID ...int64) ExcludeRequest
	Execute() (err error)
}

type executeRequest struct {
	request *api.Request
}

func (r *executeRequest) ContactTypeEmail() ExcludeRequest {
	r.request.Add("contact_type", "email")
	return r
}

func (r *executeRequest) ContactTypePhone() ExcludeRequest {
	r.request.Add("contact_type", "phone")
	return r
}

func (r *executeRequest) ListIDs(listID ...int64) ExcludeRequest {
	ids := make([]string, len(listID))
	for n, id := range listID {
		ids[n] = strconv.FormatInt(id, 10)
	}

	r.request.Add("list_ids", strings.Join(ids, ","))

	return r
}

func (r *executeRequest) Execute() (err error) {
	return r.request.Execute("exclude", nil)
}

func Exclude(request *api.Request, contact string) ExcludeRequest {
	request.Add("contact", contact)

	return &executeRequest{
		request: request,
	}
}
