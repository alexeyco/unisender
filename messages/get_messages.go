package messages

import (
	"strconv"
	"time"

	"github.com/alexeyco/unisender/api"
)

// GetMessagesResult message record.
type GetMessagesResult struct {
	ID              int64  `json:"id"`
	SubUserLogin    string `json:"sub_user_login"`
	ListID          int64  `json:"list_id"`
	SegmentID       int64  `json:"segment_id"`
	ServiceType     string `json:"service_type"`
	ActualVersionID int64  `json:"active_version_id"`
	Lang            string `json:"lang_code"`
	SenderName      string `json:"sender_name"`
	SenderEmail     string `json:"sender_email"`
	SMSFrom         string `json:"sms_from"`
	Subject         string `json:"subject"`
	Body            string `json:"body"`
	MessageFormat   string `json:"message_format"`
}

// GetMessagesRequest request to get the list of letters.
type GetMessagesRequest struct {
	request *api.Request
}

// From sets time of the message creation from which messages are to be displayed.
func (r *GetMessagesRequest) From(from time.Time) *GetMessagesRequest {
	r.request.Add("date_from", from.UTC().Format("2006-01-02 15:04:05"))
	return r
}

// To sets time of the message creation until which messages are to be displayed.
func (r *GetMessagesRequest) To(to time.Time) *GetMessagesRequest {
	r.request.Add("date_to", to.UTC().Format("2006-01-02 15:04:05"))
	return r
}

// Limit sets number of records in the response.
func (r *GetMessagesRequest) Limit(limit int) *GetMessagesRequest {
	r.request.Add("limit", strconv.Itoa(limit))
	return r
}

// Offset sets offset from which position the selection is to be started.
func (r *GetMessagesRequest) Offset(offset int) *GetMessagesRequest {
	r.request.Add("offset", strconv.Itoa(offset))
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *GetMessagesRequest) Execute() (res []GetMessagesResult, err error) {
	err = r.request.Execute("getMessages", &res)
	return
}

// GetMessages returns request to get the list of letters.
func GetMessages(request *api.Request) *GetMessagesRequest {
	return &GetMessagesRequest{
		request: request,
	}
}
