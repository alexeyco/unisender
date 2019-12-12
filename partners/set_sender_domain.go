package partners

import "github.com/alexeyco/unisender/api"

// SetSenderDomainResult result of setSenderDomain request.
type SetSenderDomainResult struct {
	DKIM string `json:"dkim"`
}

// SetSenderDomainRequest request to register the domain in the list.
type SetSenderDomainRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *SetSenderDomainRequest) Execute() (res *SetSenderDomainResult, err error) {
	var result SetSenderDomainResult
	if err = r.request.Execute("setSenderDomain", &result); err != nil {
		return
	}

	res = &result

	return
}

// SetSenderDomain register the domain in the list for authentication and generate a dkim key for it.
func SetSenderDomain(request *api.Request, login, domain string) *SetSenderDomainRequest {
	request.Add("username", login).
		Add("domain", domain)

	return &SetSenderDomainRequest{
		request: request,
	}
}
