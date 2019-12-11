package messages

import (
	"strconv"
	"time"

	"github.com/alexeyco/unisender/api"
)

// GetMessageResult information about SMS or email message.
type GetMessageResult struct {
	ID              int64     `json:"id"`
	SubUserLogin    string    `json:"sub_user_login"`
	SenderEmail     string    `json:"sender_email"`
	SenderName      string    `json:"sender_name"`
	Sender          string    `json:"sender"`
	Subject         string    `json:"subject"`
	Body            string    `json:"body"`
	BodyText        string    `json:"text_body"`
	ListID          int64     `json:"list_id"`
	LastUpdate      time.Time `json:"last_update"`
	ServiceType     string    `json:"service_type"`
	Lang            string    `json:"lang_code"`
	ActualVersionID int64     `json:"active_version_id"`
	MessageFormat   string    `json:"message_format"`
	WrapType        string    `json:"wrap_type"`
	ImagesBehavior  string    `json:"images_behavior"`
}

// GetMessageRequest request to get information about SMS or email message.
type GetMessageRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetMessageRequest) Execute() (res *GetMessageResult, err error) {
	var result GetMessageResult
	if err = r.request.Execute("getMessage", &result); err != nil {
		return
	}

	res = &result

	return
}

// GetMessage returns request to get information about SMS or email message.
func GetMessage(request *api.Request, messageID int64) *GetMessageRequest {
	request.Add("id", strconv.FormatInt(messageID, 10))

	return &GetMessageRequest{
		request: request,
	}
}
