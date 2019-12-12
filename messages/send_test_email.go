package messages

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// SendTestEmailRequest request to send a test email message.
type SendTestEmailRequest struct {
	request *api.Request
}

// To sets recipient email.
func (r *SendTestEmailRequest) To(email string) *SendTestEmailRequest {
	r.request.Add("email", email)
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *SendTestEmailRequest) Execute() (err error) {
	err = r.request.Execute("sendTestEmail", nil)
	return
}

// SendTestEmail returns request to send a test email message.
func SendTestEmail(request *api.Request, emailID int64) *SendTestEmailRequest {
	request.Add("id", strconv.FormatInt(emailID, 10))

	return &SendTestEmailRequest{
		request: request,
	}
}
