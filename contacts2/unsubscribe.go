package contacts2

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// UnsubscribeRequest request that unsubscribes the contact email or phone number from one or several lists.
type UnsubscribeRequest struct {
	request *api.Request
}

// ContactTypeEmail sets contact type to email.
func (r *UnsubscribeRequest) ContactTypeEmail() *UnsubscribeRequest {
	r.request.Add("contact_type", "email")
	return r
}

// ContactTypePhone sets contact type to phone.
func (r *UnsubscribeRequest) ContactTypePhone() *UnsubscribeRequest {
	r.request.Add("contact_type", "phone")
	return r
}

// ListIDs sets contact lists IDs from which contacts are being unsubscribed. If it is not specified,
// contacts will be unsubscribed from all lists. Codes of the lists can be obtained by calling the getLists method.
// They match the codes used in the subscription form.
//
// See https://www.unisender.com/en/support/api/partners/getlists/
func (r *UnsubscribeRequest) ListIDs(listID ...int64) *UnsubscribeRequest {
	ids := make([]string, len(listID))
	for n, id := range listID {
		ids[n] = strconv.FormatInt(id, 10)
	}

	r.request.Add("list_ids", strings.Join(ids, ","))

	return r
}

// Execute sends request to UniSender API and returns result.
func (r *UnsubscribeRequest) Execute() (err error) {
	return r.request.Execute("unsubscribe", nil)
}

// Unsubscribe returns request that unsubscribes the contact email or phone number from one or several lists.
// In contrast to the exclude method, it does not exclude a contact from the lists, but marks the contact
// as "unsubscribed". It is impossible to restore the «active» status through API – it is only the contact
// who can do this by clicking on the activation link in the letter.
//
// See https://www.unisender.com/en/support/api/partners/unsubscribe/
func Unsubscribe(request *api.Request, contact string) *UnsubscribeRequest {
	request.Add("contact", contact)

	return &UnsubscribeRequest{
		request: request,
	}
}
