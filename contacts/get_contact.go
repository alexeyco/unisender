package contacts

import "github.com/alexeyco/unisender/api"

// See https://www.unisender.com/en/support/api/contacts/getcontact/

type GetContactRequest interface {
	IncludeLists() GetContactRequest
	IncludeFields() GetContactRequest
	IncludeDetails() GetContactRequest
	Execute() (contacts []Person, err error)
}

type getContactRequest struct {
	request *api.Request
}

func (r *getContactRequest) IncludeLists() GetContactRequest {
	r.request.Add("include_lists", "1")
	return r
}

func (r *getContactRequest) IncludeFields() GetContactRequest {
	r.request.Add("include_fields", "1")
	return r
}

func (r *getContactRequest) IncludeDetails() GetContactRequest {
	r.request.Add("include_details", "1")
	return r
}

func (r *getContactRequest) Execute() (contacts []Person, err error) {
	err = r.request.Execute("getContact", &contacts)
	return
}

func GetContact(request *api.Request, email string) GetContactRequest {
	request.Add("email", email)

	return &getContactRequest{
		request: request,
	}
}
