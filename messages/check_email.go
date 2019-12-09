package messages

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// CheckEmailResponseStatus email status info.
type CheckEmailResponseStatus struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

// CheckEmailResponse request checkEmail response.
type CheckEmailResponse struct {
	Statuses []CheckEmailResponseStatus
}

// CheckEmailRequest to check the delivery status of emails sent using the sendEmail method.
type CheckEmailRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *CheckEmailRequest) Execute() (res *CheckEmailResponse, err error) {
	var response CheckEmailResponse
	if err = r.request.Execute("checkEmail", &response); err != nil {
		return
	}

	res = &response

	return
}

// CheckEmail returns request to check the delivery status of emails sent using the sendEmail method.
//
// See: https://www.unisender.com/en/support/api/messages/checkemail/
func CheckEmail(request *api.Request, emailIDs ...int64) *CheckEmailRequest {
	ids := make([]string, len(emailIDs))
	for n, id := range emailIDs {
		ids[n] = strconv.FormatInt(id, 10)
	}

	request.Add("email_id", strings.Join(ids, ","))

	return &CheckEmailRequest{
		request: request,
	}
}
