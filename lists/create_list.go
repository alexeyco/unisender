package lists

import "github.com/alexeyco/unisender/api"

// CreateListResult response of createList request.
type CreateListResult struct {
	ID int64 `json:"id"`
}

// CreateListRequest request to create a new contact list.
type CreateListRequest struct {
	request *api.Request
}

// BeforeSubscribeUrl sets the URL for redirect to the pre-subscription page. Usually this page shows a message
// that the contact should follow the confirmation link to activate the subscription. You can add substitution fields
// to this URL — for example, you can identify a contact by the email address by substituting an email here —
// or by the contact code in your database by saving the code in an additional field and substituting it into this URL.
//
// See https://www.unisender.com/en/support/campaigns/letter/substitutes/
func (r *CreateListRequest) BeforeSubscribeUrl(u string) *CreateListRequest {
	r.request.Add("before_subscribe_url", u)
	return r
}

// AfterSubscribeUrl sets the URL for redirect to the post-subscription page. Usually this page shows a message
// that the subscription has been completed successfully. You can add substitution fields to this URL — for example,
// you can identify a contact by the email address by substituting an email here — or by the contact code
// in your database by saving the code in an additional field and substituting it into this URL.
//
// See https://www.unisender.com/en/support/campaigns/letter/substitutes/
func (r *CreateListRequest) AfterSubscribeUrl(u string) *CreateListRequest {
	r.request.Add("after_subscribe_url", u)
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *CreateListRequest) Execute() (listID int64, err error) {
	var result CreateListResult
	if err = r.request.Execute("createList", &result); err != nil {
		return
	}

	listID = result.ID

	return
}

// CreateList returns request to create a new contact list.
func CreateList(request *api.Request, title string) *CreateListRequest {
	request.Add("title", title)

	return &CreateListRequest{
		request: request,
	}
}
