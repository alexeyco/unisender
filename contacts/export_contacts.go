package contacts

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// ExportContactsResult exportContacts request result.
type ExportContactsResult struct {
	TaskUUID string `json:"task_uuid"`
	Status   string `json:"status"`
}

// ExportContactsRequest request, that exports of contact data from UniSender.
type ExportContactsRequest struct {
	request *api.Request
}

// NotifyUrl sets the URL to which the response will be sent after the report is generated.
func (r *ExportContactsRequest) NotifyUrl(notifyUrl string) *ExportContactsRequest {
	r.request.Add("notify_url", notifyUrl)
	return r
}

// ListID optional export list code. If it is not specified, all lists will be exported. Codes of the lists
// can be obtained by calling the getLists method. Warning: if the list is specified, those email addresses
// that are not included in the list, but contacts have phone numbers included in this list, will be included
// in the upload result as well.
//
// https://www.unisender.com/en/support/api/partners/getlists/
func (r *ExportContactsRequest) ListID(listID int64) *ExportContactsRequest {
	r.request.Add("list_id", strconv.FormatInt(listID, 10))
	return r
}

// FieldNames an array of field names to be exported. If absent, then all possible fields are exported.
// The transfer method using HTTP: field_names[]=1&field_names[]=2. The names of the system fields do not depend
// on the language, the names of the user fields are taken from the substitution names.
func (r *ExportContactsRequest) FieldNames(fieldNames ...string) *ExportContactsRequest {
	for _, fieldName := range fieldNames {
		r.request.Add("field_names[]", fieldName)
	}

	return r
}

// Email sets email address. If this parameter is specified, the result will contain only one contact
// with such email address.
func (r *ExportContactsRequest) Email(email string) *ExportContactsRequest {
	r.request.Add("email", email)
	return r
}

// Phone sets phone number. If this parameter is specified, the result will contain only one contact
// with such phone number.
func (r *ExportContactsRequest) Phone(phone string) *ExportContactsRequest {
	r.request.Add("phone", phone)
	return r
}

// Tag sets tag parameter. If this parameter is specified, the search will take into account only contacts
// that have such tag.
func (r *ExportContactsRequest) Tag(tag string) *ExportContactsRequest {
	r.request.Add("tag", tag)
	return r
}

// EmailStatusNew if used, the result will contain only contacts with "new" email address status.
func (r *ExportContactsRequest) EmailStatusNew() *ExportContactsRequest {
	r.request.Add("email_status", "new")
	return r
}

// EmailStatusInvited if used, the result will contain only contacts with "invited" email address status.
func (r *ExportContactsRequest) EmailStatusInvited() *ExportContactsRequest {
	r.request.Add("email_status", "invited")
	return r
}

// EmailStatusActive if used, the result will contain only contacts with "active" email address status.
func (r *ExportContactsRequest) EmailStatusActive() *ExportContactsRequest {
	r.request.Add("email_status", "active")
	return r
}

// EmailStatusInactive if used, the result will contain only contacts with "inactive" email address status.
func (r *ExportContactsRequest) EmailStatusInactive() *ExportContactsRequest {
	r.request.Add("email_status", "inactive")
	return r
}

// EmailStatusUnsubscribed if used, the result will contain only contacts with "ubsubscribed" email address status.
func (r *ExportContactsRequest) EmailStatusUnsubscribed() *ExportContactsRequest {
	r.request.Add("email_status", "unsubscribed")
	return r
}

// EmailStatusBlocked if used, the result will contain only contacts with "blocked" email address status.
func (r *ExportContactsRequest) EmailStatusBlocked() *ExportContactsRequest {
	r.request.Add("email_status", "blocked")
	return r
}

// EmailStatusActivationRequested if used, the result will contain only contacts with "activation_requested" email address status.
func (r *ExportContactsRequest) EmailStatusActivationRequested() *ExportContactsRequest {
	r.request.Add("email_status", "activation_requested")
	return r
}

// PhoneStatusNew if used, if used, the result will contain only contacts with "new" phone number status.
func (r *ExportContactsRequest) PhoneStatusNew() *ExportContactsRequest {
	r.request.Add("phone_status", "new")
	return r
}

// PhoneStatusActive if used, if used, the result will contain only contacts with "active" phone number status.
func (r *ExportContactsRequest) PhoneStatusActive() *ExportContactsRequest {
	r.request.Add("phone_status", "active")
	return r
}

// PhoneStatusInactive if used, if used, the result will contain only contacts with "inactive" phone number status.
func (r *ExportContactsRequest) PhoneStatusInactive() *ExportContactsRequest {
	r.request.Add("phone_status", "inactive")
	return r
}

// PhoneStatusUnsubscribed if used, if used, the result will contain only contacts with "ubsubscribed" phone number status.
func (r *ExportContactsRequest) PhoneStatusUnsubscribed() *ExportContactsRequest {
	r.request.Add("phone_status", "unsubscribed")
	return r
}

// PhoneStatusBlocked if used, if used, the result will contain only contacts with "blocked" phone number status.
func (r *ExportContactsRequest) PhoneStatusBlocked() *ExportContactsRequest {
	r.request.Add("phone_status", "blocked")
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *ExportContactsRequest) Execute() (res *ExportContactsResult, err error) {
	var response ExportContactsResult
	if err = r.request.Execute("exportContacts", &response); err != nil {
		return
	}

	res = &response

	return
}

// ExportContacts returns request, that exports of contact data from UniSender.
//
// See https://www.unisender.com/en/support/api/contacts/exportcontacts/
func ExportContacts(request *api.Request) *ExportContactsRequest {
	return &ExportContactsRequest{
		request: request,
	}
}
