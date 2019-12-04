package lists

import (
	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/partners/getlists/

type GetListsRequest interface {
	Execute() (lists []List, err error)
}

type GetListsResponse []List

type getListsRequest struct {
	request *api.Request
}

func (r *getListsRequest) Execute() (lists []List, err error) {
	err = r.request.Execute("getLists", &lists)
	return
}

func GetLists(request *api.Request) GetListsRequest {
	return &getListsRequest{
		request: request,
	}
}
