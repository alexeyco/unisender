package messages

import (
	"strconv"
	"time"

	"github.com/alexeyco/unisender/api"
)

// ListMessagesResult message record.
type ListMessagesResult struct {
	ID            int64  `json:"id"`
	SubUserLogin  string `json:"sub_user_login"`
	ListID        int64  `json:"list_id"`
	SegmentID     int64  `json:"segment_id"`
	ServiceType   string `json:"service_type"`
	Lang          string `json:"lang_code"`
	SenderName    string `json:"sender_name"`
	SenderEmail   string `json:"sender_email"`
	Subject       string `json:"subject"`
	MessageFormat string `json:"message_format"`
}

// ListMessagesRequest request to get the list of letters.
type ListMessagesRequest struct {
	request *api.Request
}

// From sets time of the message creation from which messages are to be displayed.
func (r *ListMessagesRequest) From(from time.Time) *ListMessagesRequest {
	r.request.Add("date_from", from.UTC().Format("2006-01-02 15:04:05"))
	return r
}

// To sets time of the message creation until which messages are to be displayed.
func (r *ListMessagesRequest) To(to time.Time) *ListMessagesRequest {
	r.request.Add("date_to", to.UTC().Format("2006-01-02 15:04:05"))
	return r
}

// Limit sets number of records in the response.
func (r *ListMessagesRequest) Limit(limit int) *ListMessagesRequest {
	r.request.Add("limit", strconv.Itoa(limit))
	return r
}

// Offset sets offset from which position the selection is to be started.
func (r *ListMessagesRequest) Offset(offset int) *ListMessagesRequest {
	r.request.Add("offset", strconv.Itoa(offset))
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *ListMessagesRequest) Execute() (res []ListMessagesResult, err error) {
	err = r.request.Execute("listMessages", &res)
	return
}

// ListMessages returns request to get the list of letters but without body.
func ListMessages(request *api.Request) *ListMessagesRequest {
	return &ListMessagesRequest{
		request: request,
	}
}
