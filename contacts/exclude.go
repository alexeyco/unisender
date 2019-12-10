package contacts

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// ExcludeRequest request, that excludes the contact’s email or phone number from one or several lists.
type ExcludeRequest struct {
	request *api.Request
}

// ContactTypeEmail sets contact type to email.
func (r *ExcludeRequest) ContactTypeEmail() *ExcludeRequest {
	r.request.Add("contact_type", "email")
	return r
}

// ContactTypePhone sets contact type to phone.
func (r *ExcludeRequest) ContactTypePhone() *ExcludeRequest {
	r.request.Add("contact_type", "phone")
	return r
}

// ListIDs sets contact lists IDs from which contacts are being excluded. If it is not specified,
// contacts will be excluded from all lists. Codes of the lists can be obtained by calling the getLists method.
// They match the codes used in the subscription form.
//
// See https://www.unisender.com/en/support/api/partners/getlists/
func (r *ExcludeRequest) ListIDs(listID ...int64) *ExcludeRequest {
	ids := make([]string, len(listID))
	for n, id := range listID {
		ids[n] = strconv.FormatInt(id, 10)
	}

	r.request.Add("list_ids", strings.Join(ids, ","))

	return r
}

// Execute sends request to UniSender API and returns result.
func (r *ExcludeRequest) Execute() (err error) {
	return r.request.Execute("exclude", nil)
}

// Exclude returns request, that excludes the contact’s email or phone number from one or several lists.
func Exclude(request *api.Request, contact string) *ExcludeRequest {
	request.Add("contact", contact)

	return &ExcludeRequest{
		request: request,
	}
}
