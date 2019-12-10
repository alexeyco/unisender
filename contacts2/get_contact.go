package contacts2

import "github.com/alexeyco/unisender/api"

// GetContactRequest request to getting information about a contact (one contact only).
type GetContactRequest struct {
	request *api.Request
}

// IncludeLists displays information about the lists to which the contact has been added.
func (r *GetContactRequest) IncludeLists() *GetContactRequest {
	r.request.Add("include_lists", "1")
	return r
}

// IncludeFields displays information about additional contact fields.
func (r *GetContactRequest) IncludeFields() *GetContactRequest {
	r.request.Add("include_fields", "1")
	return r
}

// IncludeDetails displays additional information about the contact.
func (r *GetContactRequest) IncludeDetails() *GetContactRequest {
	r.request.Add("include_details", "1")
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *GetContactRequest) Execute() (person *Person, err error) {
	var p Person
	if err = r.request.Execute("getContact", &p); err != nil {
		return
	}

	person = &p

	return
}

// GetContact returns request to getting information about a contact (one contact only).
//
// See https://www.unisender.com/en/support/api/contacts/getcontact/
func GetContact(request *api.Request, email string) *GetContactRequest {
	request.Add("email", email)

	return &GetContactRequest{
		request: request,
	}
}
