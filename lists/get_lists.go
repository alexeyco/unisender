package lists

import "github.com/alexeyco/unisender/api"

// GetListsResult response of createList request.
type GetListsResult struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// GetListsRequest request to get the list of all available campaign lists.
type GetListsRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetListsRequest) Execute() (lists []GetListsResult, err error) {
	err = r.request.Execute("getLists", &lists)
	return
}

// GetLists returns request to get the list of all available campaign lists.
func GetLists(request *api.Request) *GetListsRequest {
	return &GetListsRequest{
		request: request,
	}
}
