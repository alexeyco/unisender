package contacts

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// DeleteTagRequest request to delete a user tag.
type DeleteTagRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *DeleteTagRequest) Execute() (err error) {
	err = r.request.Execute("deleteTag", nil)
	return
}

// DeleteTag returns request to delete a user tag.
func DeleteTag(request *api.Request, tagID int64) *DeleteTagRequest {
	request.Add("id", strconv.FormatInt(tagID, 10))

	return &DeleteTagRequest{
		request: request,
	}
}
