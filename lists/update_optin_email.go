package lists

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

type UpdateOptInEmailRequest struct {
	request *api.Request
}

// SenderName sets sender name.
func (r *UpdateOptInEmailRequest) SenderName(senderName string) *UpdateOptInEmailRequest {
	r.request.Add("sender_name", senderName)
	return r
}

// SenderEmail sets sender email.
func (r *UpdateOptInEmailRequest) SenderEmail(senderEmail string) *UpdateOptInEmailRequest {
	r.request.Add("sender_email", senderEmail)
	return r
}

// Subject sets email subject.
func (r *UpdateOptInEmailRequest) Subject(subject string) *UpdateOptInEmailRequest {
	r.request.Add("subject", subject)
	return r
}

// Body sets email body.
func (r *UpdateOptInEmailRequest) Body(body string) *UpdateOptInEmailRequest {
	r.request.Add("body", body)
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *UpdateOptInEmailRequest) Execute() error {
	return r.request.Execute("updateOptInEmail", nil)
}

// UpdateOptInEmail returns request to update OptInEmail text.
func UpdateOptInEmail(request *api.Request, listID int64) *UpdateOptInEmailRequest {
	request.Add("list_id", strconv.FormatInt(listID, 10))

	return &UpdateOptInEmailRequest{
		request: request,
	}
}
