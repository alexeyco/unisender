package messages

import (
	"strconv"
	"time"

	"github.com/alexeyco/unisender/api"
)

// ListTemplatesResult message record.
type ListTemplatesResult struct {
	ID                    int64       `json:"id"`
	SubUserLogin          string      `json:"sub_user_login"`
	Title                 string      `json:"title"`
	Description           string      `json:"description"`
	Lang                  string      `json:"lang_code"`
	Subject               string      `json:"subject"`
	Attachments           interface{} `json:"attachments"`
	ScreenshotURL         string      `json:"screenshot_url"`
	Created               time.Time   `json:"created"`
	MessageFormat         string      `json:"message_format"`
	Type                  string      `json:"type"`
	FullsizeScreenshotURL string      `json:"fullsize_screenshot_url"`
}

// ListTemplatesRequest request to get the list of letters.
type ListTemplatesRequest struct {
	request *api.Request
}

// TypeUser sets type to user.
func (r *ListTemplatesRequest) TypeUser() *ListTemplatesRequest {
	r.request.Add("type", "user")
	return r
}

// TypeSystem sets type to system.
func (r *ListTemplatesRequest) TypeSystem() *ListTemplatesRequest {
	r.request.Add("type", "system")
	return r
}

// From sets time of the message creation from which messages are to be displayed.
func (r *ListTemplatesRequest) From(from time.Time) *ListTemplatesRequest {
	r.request.Add("date_from", from.UTC().Format("2006-01-02 15:04:05"))
	return r
}

// To sets time of the message creation until which messages are to be displayed.
func (r *ListTemplatesRequest) To(to time.Time) *ListTemplatesRequest {
	r.request.Add("date_to", to.UTC().Format("2006-01-02 15:04:05"))
	return r
}

// Limit sets number of records in the response.
func (r *ListTemplatesRequest) Limit(limit int) *ListTemplatesRequest {
	r.request.Add("limit", strconv.Itoa(limit))
	return r
}

// Offset sets offset from which position the selection is to be started.
func (r *ListTemplatesRequest) Offset(offset int) *ListTemplatesRequest {
	r.request.Add("offset", strconv.Itoa(offset))
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *ListTemplatesRequest) Execute() (res []ListTemplatesResult, err error) {
	err = r.request.Execute("listTemplates", &res)
	return
}

// ListTemplates returns request to get the list of templates.
func ListTemplates(request *api.Request) *ListTemplatesRequest {
	return &ListTemplatesRequest{
		request: request,
	}
}
