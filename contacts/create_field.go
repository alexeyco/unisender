package contacts

import "github.com/alexeyco/unisender/api"

// CreateFieldResult result of createField request.
type CreateFieldResult struct {
	ID int64 `json:"id"`
}

// CreateFieldRequest request to create a new user field.
type CreateFieldRequest struct {
	request *api.Request
}

// TypeString sets type to string.
func (r *CreateFieldRequest) TypeString() *CreateFieldRequest {
	r.request.Add("type", "string")
	return r
}

// TypeText sets type to text.
func (r *CreateFieldRequest) TypeText() *CreateFieldRequest {
	r.request.Add("type", "text")
	return r
}

// TypeNumber sets type to number.
func (r *CreateFieldRequest) TypeNumber() *CreateFieldRequest {
	r.request.Add("type", "number")
	return r
}

// TypeDate sets type to date.
func (r *CreateFieldRequest) TypeDate() *CreateFieldRequest {
	r.request.Add("type", "date")
	return r
}

// TypeBool sets type to bool.
func (r *CreateFieldRequest) TypeBool() *CreateFieldRequest {
	r.request.Add("type", "bool")
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *CreateFieldRequest) Execute() (fieldID int64, err error) {
	var result CreateFieldResult
	if err = r.request.Execute("createField", &result); err != nil {
		return
	}

	fieldID = result.ID

	return
}

// CreateField returns request to create a new user field, the value of which can be set for each recipient,
// and then it can be substituted in the letter.
func CreateField(request *api.Request, name string) *CreateFieldRequest {
	request.Add("name", name)

	return &CreateFieldRequest{
		request: request,
	}
}
