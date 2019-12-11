package messages

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// CreateSMSMessageResult createSMSMessage request result.
type CreateSMSMessageResult struct {
	MessageID int64 `json:"message_id"`
}

// CreateSMSMessageRequest request to create an sms without sending it.
type CreateSMSMessageRequest struct {
	request *api.Request
}

// Body sets message body.
func (r *CreateSMSMessageRequest) Body(body string) *CreateSMSMessageRequest {
	r.request.Add("body", body)
	return r
}

// ListID sets message list ID.
func (r *CreateSMSMessageRequest) ListID(listID int64) *CreateSMSMessageRequest {
	r.request.Add("list_id", strconv.FormatInt(listID, 10))
	return r
}

// Body sets message tag.
func (r *CreateSMSMessageRequest) Tag(tag string) *CreateSMSMessageRequest {
	r.request.Add("tag", tag)
	return r
}

// Categories sets message categories.
func (r *CreateSMSMessageRequest) Categories(categories ...string) *CreateSMSMessageRequest {
	r.request.Add("categories", strings.Join(categories, ","))
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *CreateSMSMessageRequest) Execute() (messageID int64, err error) {
	var result CreateSMSMessageResult
	if err = r.request.Execute("createSmsMessage", &result); err != nil {
		return
	}

	messageID = result.MessageID

	return
}

// CreateSMSMessage returns request to create an sms without sending it.
func CreateSMSMessage(request *api.Request, sender string) *CreateSMSMessageRequest {
	request.Add("sender", sender)

	return &CreateSMSMessageRequest{
		request: request,
	}
}
