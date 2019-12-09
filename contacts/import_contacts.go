package contacts

import (
	"fmt"

	"github.com/alexeyco/unisender/api"
)

// ImportContactsResponseLogMessage importContacts log messages.
type ImportContactsResponseLogMessage struct {
	Index   int    `json:"index"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ImportContactsResponse response of importContacts request.
type ImportContactsResponse struct {
	Total     int                                `json:"total"`
	Inserted  int                                `json:"inserted"`
	Updated   int                                `json:"updated"`
	Deleted   int                                `json:"deleted"`
	NewEmails int                                `json:"new_emails"`
	Invalid   int                                `json:"invalid"`
	Log       []ImportContactsResponseLogMessage `json:"log"`
}

// ImportContactsRequest request to bulk import of contacts.
type ImportContactsRequest struct {
	request    *api.Request
	collection *ImportContactsCollection
	fieldNames []string
}

// OverwriteTags if used, contacts tags will be replaced.
func (r *ImportContactsRequest) OverwriteTags() *ImportContactsRequest {
	r.request.Add("overwrite_tags", "1")
	return r
}

// OverwriteLists if used, contacts lists will be replaced.
func (r *ImportContactsRequest) OverwriteLists() *ImportContactsRequest {
	r.request.Add("overwrite_lists", "1")
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *ImportContactsRequest) Execute() (res *ImportContactsResponse, err error) {
	for n, fieldName := range r.collection.FieldNames() {
		r.request.Add(fmt.Sprintf("field_names[%d]", n), fieldName)
	}

	data := r.collection.Data()
	for row, c := range data {
		for col, val := range c {
			r.request.Add(fmt.Sprintf("data[%d][%d]", row, col), val)
		}
	}

	var result ImportContactsResponse
	if err = r.request.Execute("importContacts", &result); err != nil {
		return
	}

	res = &result

	return
}

// ImportContacts returns request to bulk import of contacts. It can also be used for periodic synchronization
// with the contact database stored on your own server (see also the description of the exportContacts method).
// You can import data of no more than 500 contacts per call. Larger lists must be imported in a few calls.
//
// If there are new addresses among the signed e-mail addresses, then by default they receive the status "new".
//
// Technical restrictions: the maximum number of user fields is 50. The timeout per call is 30 seconds from the moment
// the request is completely transmitted to the server. If no response is received after the timeout,
// then it is recommended to make up to two retries, and if there is no answer again, then contact technical support.
//
// See https://www.unisender.com/en/support/api/contacts/importcontacts/
func ImportContacts(request *api.Request, collection *ImportContactsCollection) *ImportContactsRequest {
	return &ImportContactsRequest{
		request:    request,
		collection: collection,
	}
}
