package contacts

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// IsContactInListRequest request to check whether the contact is in the specified user lists.
type IsContactInListRequest struct {
	request *api.Request
}

// ConditionOr if used, request checks if the contact is in at least one of the specified lists.
func (r *IsContactInListRequest) ConditionOr() *IsContactInListRequest {
	r.request.Add("condition", "or")
	return r
}

// ConditionAnd if used, request checks if the contact is in all specified lists.
func (r *IsContactInListRequest) ConditionAnd() *IsContactInListRequest {
	r.request.Add("condition", "and")
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *IsContactInListRequest) Execute() (res bool, err error) {
	err = r.request.Execute("isContactInList", &res)
	return
}

// IsContactInList returns request to check whether the contact is in the specified user lists.
func IsContactInList(request *api.Request, email string, listIDs ...int64) *IsContactInListRequest {
	ids := make([]string, len(listIDs))
	for n, id := range listIDs {
		ids[n] = strconv.FormatInt(id, 10)
	}

	request.Add("email", email).
		Add("list_ids", strings.Join(ids, ","))

	return &IsContactInListRequest{
		request: request,
	}
}
