package contacts

import (
	"fmt"

	"github.com/alexeyco/unisender/api"
)

// See https://www.unisender.com/en/support/api/contacts/exportcontacts/

type ExportContactsRequest interface {
	NotifyUrl(notifyUrl string) ExportContactsRequest
	ListID(listID int64) ExportContactsRequest
	FieldNames(fieldNames ...string) ExportContactsRequest
	Email(email string) ExportContactsRequest
	Phone(phone string) ExportContactsRequest
	Tag(tag string) ExportContactsRequest
	EmailStatusNew() ExportContactsRequest
	EmailStatusInvited() ExportContactsRequest
	EmailStatusActive() ExportContactsRequest
	EmailStatusInactive() ExportContactsRequest
	EmailStatusUnsubscribed() ExportContactsRequest
	EmailStatusBlocked() ExportContactsRequest
	EmailStatusActivationRequested() ExportContactsRequest
	PhoneStatusNew() ExportContactsRequest
	PhoneStatusActive() ExportContactsRequest
	PhoneStatusInactive() ExportContactsRequest
	PhoneStatusUnsubscribed() ExportContactsRequest
	PhoneStatusBlocked() ExportContactsRequest
	Execute() (res *ExportContactResult, err error)
}

type ExportContactResult struct {
	TaskUUID string `json:"task_uuid"`
	Status   string `json:"status"`
}

type exportContactsRequest struct {
	request *api.Request
}

func (r *exportContactsRequest) NotifyUrl(notifyUrl string) ExportContactsRequest {
	r.request.Add("notify_url", notifyUrl)
	return r
}

func (r *exportContactsRequest) ListID(listID int64) ExportContactsRequest {
	r.request.Add("list_id", fmt.Sprintf("%d", listID))
	return r
}

func (r *exportContactsRequest) FieldNames(fieldNames ...string) ExportContactsRequest {
	for _, fieldName := range fieldNames {
		r.request.Add("field_names[]", fieldName)
	}

	return r
}

func (r *exportContactsRequest) Email(email string) ExportContactsRequest {
	r.request.Add("email", email)
	return r
}

func (r *exportContactsRequest) Phone(phone string) ExportContactsRequest {
	r.request.Add("phone", phone)
	return r
}

func (r *exportContactsRequest) Tag(tag string) ExportContactsRequest {
	r.request.Add("tag", tag)
	return r
}

func (r *exportContactsRequest) EmailStatusNew() ExportContactsRequest {
	r.request.Add("email_status", "new")
	return r
}

func (r *exportContactsRequest) EmailStatusInvited() ExportContactsRequest {
	r.request.Add("email_status", "invited")
	return r
}

func (r *exportContactsRequest) EmailStatusActive() ExportContactsRequest {
	r.request.Add("email_status", "active")
	return r
}

func (r *exportContactsRequest) EmailStatusInactive() ExportContactsRequest {
	r.request.Add("email_status", "inactive")
	return r
}

func (r *exportContactsRequest) EmailStatusUnsubscribed() ExportContactsRequest {
	r.request.Add("email_status", "unsubscribed")
	return r
}

func (r *exportContactsRequest) EmailStatusBlocked() ExportContactsRequest {
	r.request.Add("email_status", "blocked")
	return r
}

func (r *exportContactsRequest) EmailStatusActivationRequested() ExportContactsRequest {
	r.request.Add("email_status", "activation_requested")
	return r
}

func (r *exportContactsRequest) PhoneStatusNew() ExportContactsRequest {
	r.request.Add("phone_status", "new")
	return r
}

func (r *exportContactsRequest) PhoneStatusActive() ExportContactsRequest {
	r.request.Add("phone_status", "active")
	return r
}

func (r *exportContactsRequest) PhoneStatusInactive() ExportContactsRequest {
	r.request.Add("phone_status", "inactive")
	return r
}

func (r *exportContactsRequest) PhoneStatusUnsubscribed() ExportContactsRequest {
	r.request.Add("phone_status", "unsubscribed")
	return r
}

func (r *exportContactsRequest) PhoneStatusBlocked() ExportContactsRequest {
	r.request.Add("phone_status", "blocked")
	return r
}

func (r *exportContactsRequest) Execute() (res *ExportContactResult, err error) {
	var response ExportContactResult
	if err = r.request.Execute("exportContacts", &response); err != nil {
		return
	}

	res = &response

	return
}

func ExportContacts(request *api.Request) ExportContactsRequest {
	return &exportContactsRequest{
		request: request,
	}
}
