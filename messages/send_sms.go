package messages

import (
	"github.com/alexeyco/unisender/api"
	"strings"
)

// SendSMSResult result sendSms request.
type SendSMSResult struct {
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
	SMSID    int64   `json:"sms_id"`
}

// SendSMSRequest
type SendSMSRequest struct {
	request *api.Request
}

// Sender sets SMS message sender.
func (r *SendSMSRequest) Sender(sender string) *SendSMSRequest {
	r.request.Add("sender", sender)
	return r
}

// Text sets SMS text.
func (r *SendSMSRequest) Text(text string) *SendSMSRequest {
	r.request.Add("text", text)
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *SendSMSRequest) Execute() (res *SendSMSResult, err error) {
	var result SendSMSResult
	if err = r.request.Execute("sendSms", &result); err != nil {
		return
	}

	res = &result

	return
}

// SendSMS returns request for sending the one SMS to one or several recipients.
func SendSMS(request *api.Request, phone ...string) *SendSMSRequest {
	request.Add("phone", strings.Join(phone, ","))

	return &SendSMSRequest{
		request: request,
	}
}
