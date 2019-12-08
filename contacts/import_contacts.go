package contacts

import (
	"fmt"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/contacts/importcontacts/

type ImportContactsRequest interface {
	FieldNames(fieldNames ...string) ImportContactsRequest
	OverwriteTags() ImportContactsRequest
	OverwriteLists() ImportContactsRequest
	Execute() (res *ImportContactsResponse, err error)
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
	fieldNames []string
}

func (r *importContactsRequest) FieldNames(fieldNames ...string) ImportContactsRequest {
	r.request.Add("field_names", strings.Join(fieldNames, ","))
	r.fieldNames = fieldNames

	return r
}

func (r *importContactsRequest) OverwriteTags() ImportContactsRequest {
	r.request.Add("overwrite_tags", "1")
	return r
}

func (r *importContactsRequest) OverwriteLists() ImportContactsRequest {
	r.request.Add("overwrite_lists", "1")
	return r
}

func (r *importContactsRequest) Execute() (res *ImportContactsResponse, err error) {
	data := r.collection.Data(r.fieldNames...)
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

func ImportContacts(request *api.Request, collection *Collection) ImportContactsRequest {
	return &importContactsRequest{
		request:    request,
		collection: collection,
	}
}
