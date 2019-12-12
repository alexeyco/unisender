package partners

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// GetSenderDomainListResultDomain domain info.
type GetSenderDomainListResultDomain struct {
	Domain string `json:"Domain"`
	Status string `json:"Status"`
	Key    string `json:"key"`
}

// GetSenderDomainListResult domain list.
type GetSenderDomainListResult struct {
	Status  string                            `json:"status"`
	Domains []GetSenderDomainListResultDomain `json:"domains"`
}

// GetSenderDomainListRequest request to get domains information.
type GetSenderDomainListRequest struct {
	request *api.Request
}

// Domain sets domain name.
func (r *GetSenderDomainListRequest) Domain(domain string) *GetSenderDomainListRequest {
	r.request.Add("domain", domain)
	return r
}

// Limit sets the number of entries in the response to one request must be an integer
// in the range of 1 – 100, the default is 50.
func (r *GetSenderDomainListRequest) Limit(limit int) *GetSenderDomainListRequest {
	r.request.Add("limit", strconv.Itoa(limit))
	return r
}

// Offset shows which position to start sampling from. Values: 0 or more are accepted
// (the position of the first record starts from 0), the default is 0.
func (r *GetSenderDomainListRequest) Offset(offset int) *GetSenderDomainListRequest {
	r.request.Add("offset", strconv.Itoa(offset))
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *GetSenderDomainListRequest) Execute() (res *GetSenderDomainListResult, err error) {
	var result GetSenderDomainListResult
	if err = r.request.Execute("getSenderDomainList", &result); err != nil {
		return
	}

	res = &result

	return
}

// GetSenderDomainList returns information about domains.
func GetSenderDomainList(request *api.Request, login string) *GetSenderDomainListRequest {
	request.Add("username", login).
		Add("format", "json")

	return &GetSenderDomainListRequest{
		request: request,
	}
}
