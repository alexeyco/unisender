package contacts

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/contacts/iscontactinlist/

type IsContactInListRequest interface {
	ConditionOr() IsContactInListRequest
	ConditionAnd() IsContactInListRequest
	Execute() (res bool, err error)
}

type isContactInListRequest struct {
	request *api.Request
}

func (r *isContactInListRequest) ConditionOr() IsContactInListRequest {
	r.request.Add("condition", "or")
	return r
}

func (r *isContactInListRequest) ConditionAnd() IsContactInListRequest {
	r.request.Add("condition", "and")
	return r
}

func (r *isContactInListRequest) Execute() (res bool, err error) {
	err = r.request.Execute("isContactInList", &res)
	return
}

func IsContactInList(request *api.Request, email string, listIDs ...int64) IsContactInListRequest {
	ids := make([]string, len(listIDs))
	for n, id := range listIDs {
		ids[n] = strconv.FormatInt(id, 10)
	}

	request.Add("email", email).
		Add("list_ids", strings.Join(ids, ","))

	return &isContactInListRequest{
		request: request,
	}
}
