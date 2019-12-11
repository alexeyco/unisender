package contacts

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// UpdateFieldResult result with updated user field ID.
type UpdateFieldResult struct {
	ID int64 `json:"id"`
}

// UpdateFieldRequest request to change user field parameters.
type UpdateFieldRequest struct {
	request *api.Request
}

// PublicName sets name of the "variable for substitution" field in the personal account. If it is not used,
// an automatically generation by the "name" field will take place. Admissible characters: latin letters,
// numbers, "_" and "-". The first character must be a letter. No spaces are allowed.
func (r *UpdateFieldRequest) PublicName(publicName string) *UpdateFieldRequest {
	r.request.Add("public_name", publicName)
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *UpdateFieldRequest) Execute() (fieldID int64, err error) {
	var result UpdateFieldResult
	if err = r.request.Execute("updateField", &result); err != nil {
		return
	}

	fieldID = result.ID

	return
}

// UpdateField returns request to change user field parameters.
func UpdateField(request *api.Request, fieldID int64, name string) *UpdateFieldRequest {
	request.Add("id", strconv.FormatInt(fieldID, 10)).
		Add("name", name)

	return &UpdateFieldRequest{
		request: request,
	}
}
