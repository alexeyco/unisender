package messages

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// CheckSMSResult request checkSMS response.
type CheckSMSResult struct {
	Status string
}

// CheckSMSRequest to check the delivery status of sms sent using the sendSMS method.
type CheckSMSRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *CheckSMSRequest) Execute() (res string, err error) {
	var result CheckSMSResult
	if err = r.request.Execute("checkSms", &result); err != nil {
		return
	}

	res = result.Status

	return
}

// CheckSMS returns request to check the delivery status of sms sent using the sendSMS method.
func CheckSMS(request *api.Request, smsID int64) *CheckSMSRequest {
	request.Add("sms_id", strconv.FormatInt(smsID, 10))

	return &CheckSMSRequest{
		request: request,
	}
}
