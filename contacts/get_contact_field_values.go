package contacts

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// GetContactFieldValuesResult field values.
type GetContactFieldValuesResult struct {
	FieldValues map[int64]string `json:"fieldValues"`
}

// GetContactFieldValuesRequest request to obtain the values of the specified additional contact fields.
type GetContactFieldValuesRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetContactFieldValuesRequest) Execute() (res map[int64]string, err error) {
	var response GetContactFieldValuesResult
	if err = r.request.Execute("getContactFieldValues", &response); err != nil {
		return
	}

	res = response.FieldValues

	return
}

// GetContactFieldValues returns request to obtain the values of the specified additional contact fields.
func GetContactFieldValues(request *api.Request, email string, fieldIDs ...int64) *GetContactFieldValuesRequest {
	f := make([]string, len(fieldIDs))
	for n, fieldID := range fieldIDs {
		f[n] = strconv.FormatInt(fieldID, 10)
	}

	request.Add("email", email).
		Add("field_ids", strings.Join(f, ","))

	return &GetContactFieldValuesRequest{
		request: request,
	}
}
