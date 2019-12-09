package contacts

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// UpdateListRequest request to change campaign list properties.
type UpdateListRequest struct {
	request *api.Request
}

// BeforeSubscribeUrl sets the URL for redirect to the pre-subscription page. Usually this page shows a message
// that the contact should follow the confirmation link to activate the subscription. You can add substitution fields
// to this URL — for example, you can identify a contact by the email address by substituting an email here —
// or by the contact code in your database by saving the code in an additional field and substituting it into this URL.
//
// See https://www.unisender.com/en/support/campaigns/letter/substitutes/
func (r *UpdateListRequest) BeforeSubscribeUrl(u string) *UpdateListRequest {
	r.request.Add("before_subscribe_url", u)
	return r
}

// AfterSubscribeUrl sets the URL for redirect to the post-subscription page. Usually this page shows a message
// that the subscription has been completed successfully. You can add substitution fields to this URL — for example,
// you can identify a contact by the email address by substituting an email here — or by the contact code
// in your database by saving the code in an additional field and substituting it into this URL.
//
// See https://www.unisender.com/en/support/campaigns/letter/substitutes/
func (r *UpdateListRequest) AfterSubscribeUrl(u string) *UpdateListRequest {
	r.request.Add("after_subscribe_url", u)
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *UpdateListRequest) Execute() error {
	return r.request.Execute("updateList", nil)
}

// UpdateList returns request to change campaign list properties.
//
// See https://www.unisender.com/en/support/api/partners/updatelist/
func UpdateList(request *api.Request, listID int64, title string) *UpdateListRequest {
	request.Add("list_id", strconv.FormatInt(listID, 10)).
		Add("title", title)

	return &UpdateListRequest{
		request: request,
	}
}
