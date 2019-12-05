package contacts

import "github.com/alexeyco/unisender/api"

// See https://www.unisender.com/en/support/api/contacts/importcontacts/

type ImportContactsRequest interface {
	Execute() (err error)
}

type ImportContactsResponse struct {
	Total     int                                `json:"total"`
	Inserted  int                                `json:"inserted"`
	Updated   int                                `json:"updated"`
	Deleted   int                                `json:"deleted"`
	NewEmails int                                `json:"new_emails"`
	Invalid   int                                `json:"invalid"`
	Log       []ImportContactsResponseLogMessage `json:"log"`
}

type ImportContactsResponseLogMessage struct {
	Index   int    `json:"index"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type importContactsRequest struct {
	request    *api.Request
	collection *Collection
}

func (r *importContactsRequest) OverwriteTags() ImportContactsRequest {
	r.request.Add("overwrite_tags", "1")
	return r
}

func (r *importContactsRequest) OverwriteLists() ImportContactsRequest {
	r.request.Add("overwrite_lists", "1")
	return r
}

func (r *importContactsRequest) Execute() (err error) {
	return
}

func ImportContacts(request *api.Request, collection *Collection) ImportContactsRequest {
	return &importContactsRequest{
		request:    request,
		collection: collection,
	}
}
