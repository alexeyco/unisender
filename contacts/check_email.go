package contacts

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// CheckEmailResultStatus email status info.
type CheckEmailResultStatus struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

// CheckEmailResult request checkEmail response.
type CheckEmailResult struct {
	Statuses []CheckEmailResultStatus
}

// CheckEmailRequest to check the delivery status of emails sent using the sendEmail method.
type CheckEmailRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *CheckEmailRequest) Execute() (res *CheckEmailResult, err error) {
	var result CheckEmailResult
	if err = r.request.Execute("checkEmail", &result); err != nil {
		return
	}

	res = &result

	return
}

// CheckEmail returns request to check the delivery status of emails sent using the sendEmail method.
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
