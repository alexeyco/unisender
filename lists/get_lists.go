package lists

import (
	unisender2 "github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/partners/getlists/

type GetListsRequest interface {
	Execute() (lists []List, err error)
}

type GetListsResponse []List

type getListsRequest struct {
	request *unisender2.Request
}

func (r *getListsRequest) Execute() (lists []List, err error) {
	err = r.request.Execute("getLists", lists)
	return
}

func GetLists(request *unisender2.Request) GetListsRequest {
	return &getListsRequest{
		request: request,
	}
}
