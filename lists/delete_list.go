package lists

import (
	"fmt"

	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/partners/deletelist/

type DeleteListRequest interface {
	Execute() error
}

type deleteListRequest struct {
	request *api.Request
}

func (r *deleteListRequest) Execute() error {
	return r.request.Execute("deleteList", nil)
}

func DeleteList(request *api.Request, listID int64) DeleteListRequest {
	request.Add("list_id", fmt.Sprintf("%d", listID))

	return &deleteListRequest{
		request: request,
	}
}
