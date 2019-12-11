package contacts

import (
	"fmt"

	"github.com/alexeyco/unisender/api"
)

// ImportContactsResultLogMessage importContacts log messages.
type ImportContactsResultLogMessage struct {
	Index   int    `json:"index"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ImportContactsResultResult response of importContacts request.
type ImportContactsResultResult struct {
	Total     int                              `json:"total"`
	Inserted  int                              `json:"inserted"`
	Updated   int                              `json:"updated"`
	Deleted   int                              `json:"deleted"`
	NewEmails int                              `json:"new_emails"`
	Invalid   int                              `json:"invalid"`
	Log       []ImportContactsResultLogMessage `json:"log"`
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
func (r *ImportContactsRequest) Execute() (res *ImportContactsResultResult, err error) {
	for n, fieldName := range r.collection.FieldNames() {
		r.request.Add(fmt.Sprintf("field_names[%d]", n), fieldName)
	}

	data := r.collection.Data()
	for row, c := range data {
		for col, val := range c {
			r.request.Add(fmt.Sprintf("data[%d][%d]", row, col), val)
		}
	}

	var result ImportContactsResultResult
	if err = r.request.Execute("importContacts", &result); err != nil {
		return
	}

	res = &result

	return
}

// ImportContacts returns request to bulk import of contacts.
func ImportContacts(request *api.Request, collection *ImportContactsCollection) *ImportContactsRequest {
	return &ImportContactsRequest{
		request:    request,
		collection: collection,
	}
}
