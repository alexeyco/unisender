package contacts

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// DeleteFieldRequest request to delete a user field.
type DeleteFieldRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *DeleteFieldRequest) Execute() (err error) {
	err = r.request.Execute("deleteField", nil)
	return
}

// DeleteField returns request to delete a user field.
func DeleteField(request *api.Request, fieldID int64) *DeleteFieldRequest {
	request.Add("id", strconv.FormatInt(fieldID, 10))

	return &DeleteFieldRequest{
		request: request,
	}
}
