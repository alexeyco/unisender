package contacts

import "github.com/alexeyco/unisender/api"

// GetFieldsResult user fields list.
type GetFieldsResult struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	IsVisible int    `json:"is_visible"`
	ViewPos   int    `json:"view_pos"`
}

// GetFieldsRequest request to get the list of user fields.
type GetFieldsRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetFieldsRequest) Execute() (res []GetFieldsResult, err error) {
	err = r.request.Execute("getFields", &res)
	return
}

// GetFields returns request to get the list of user fields.
func GetFields(request *api.Request) *GetFieldsRequest {
	return &GetFieldsRequest{
		request: request,
	}
}
