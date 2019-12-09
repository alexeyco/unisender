package messages

import "github.com/alexeyco/unisender/api"

// SetSenderDomainResponse result of setSenderDomain request.
type SetSenderDomainResponse struct {
	DKIM string `json:"dkim"`
}

// SetSenderDomainRequest request to register the domain in the list.
type SetSenderDomainRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *SetSenderDomainRequest) Execute() (res *SetSenderDomainResponse, err error) {
	var response SetSenderDomainResponse
	if err = r.request.Execute("setSenderDomain", &response); err != nil {
		return
	}

	res = &response

	return
}

// SetSenderDomain register the domain in the list for authentication and generate a dkim key for it. Confirm
// the address on the domain to add the domain to the list.
//
// See: https://www.unisender.com/en/support/api/messages/setsenderdomain/
func SetSenderDomain(request *api.Request, login, domain string) *SetSenderDomainRequest {
	request.Add("username", login).
		Add("domain", domain)

	return &SetSenderDomainRequest{
		request: request,
	}
}
