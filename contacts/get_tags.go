package contacts

import "github.com/alexeyco/unisender/api"

// GetTagsResult user tag.
type GetTagsResult struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GetTagsRequest request to get user tags.
type GetTagsRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetTagsRequest) Execute() (res []GetTagsResult, err error) {
	err = r.request.Execute("getTags", &res)
	return
}

// GetTags returns request to get user tags.
func GetTags(request *api.Request) *GetTagsRequest {
	return &GetTagsRequest{
		request: request,
	}
}
