package messages2

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// GetSenderDomainListResponseDomain domain info.
type GetSenderDomainListResponseDomain struct {
	Domain string `json:"Domain"`
	Status string `json:"Status"`
	Key    string `json:"key"`
}

// GetSenderDomainListResponse domain list.
type GetSenderDomainListResponse struct {
	Status  string                              `json:"status"`
	Domains []GetSenderDomainListResponseDomain `json:"domains"`
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
func (r *GetSenderDomainListRequest) Execute() (res *GetSenderDomainListResponse, err error) {
	var response GetSenderDomainListResponse
	if err = r.request.Execute("getSenderDomainList", &response); err != nil {
		return
	}

	res = &response

	return
}

// GetSenderDomainList returns information about domains.
//
// See: https://www.unisender.com/en/support/api/messages/getsenderdomainlist/
func GetSenderDomainList(request *api.Request, login string) *GetSenderDomainListRequest {
	request.Add("username", login).
		Add("format", "json")

	return &GetSenderDomainListRequest{
		request: request,
	}
}
