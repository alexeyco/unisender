package messages

import (
	"strconv"
	"time"

	"github.com/alexeyco/unisender/api"
)

// GetTemplatesResult message record.
type GetTemplatesResult struct {
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
	Body                  string      `json:"body"`
	BodyRaw               string      `json:"raw_body"`
	FullsizeScreenshotURL string      `json:"fullsize_screenshot_url"`
}

// GetTemplatesRequest request to get the list of letters.
type GetTemplatesRequest struct {
	request *api.Request
}

// TypeUser sets type to user.
func (r *GetTemplatesRequest) TypeUser() *GetTemplatesRequest {
	r.request.Add("type", "user")
	return r
}

// TypeSystem sets type to system.
func (r *GetTemplatesRequest) TypeSystem() *GetTemplatesRequest {
	r.request.Add("type", "system")
	return r
}

// From sets time of the message creation from which messages are to be displayed.
func (r *GetTemplatesRequest) From(from time.Time) *GetTemplatesRequest {
	r.request.Add("date_from", from.UTC().Format("2006-01-02 15:04:05"))
	return r
}

// To sets time of the message creation until which messages are to be displayed.
func (r *GetTemplatesRequest) To(to time.Time) *GetTemplatesRequest {
	r.request.Add("date_to", to.UTC().Format("2006-01-02 15:04:05"))
	return r
}

// Limit sets number of records in the response.
func (r *GetTemplatesRequest) Limit(limit int) *GetTemplatesRequest {
	r.request.Add("limit", strconv.Itoa(limit))
	return r
}

// Offset sets offset from which position the selection is to be started.
func (r *GetTemplatesRequest) Offset(offset int) *GetTemplatesRequest {
	r.request.Add("offset", strconv.Itoa(offset))
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *GetTemplatesRequest) Execute() (res []GetTemplatesResult, err error) {
	err = r.request.Execute("getTemplates", &res)
	return
}

// GetTemplates returns request to get the list of templates.
func GetTemplates(request *api.Request) *GetTemplatesRequest {
	return &GetTemplatesRequest{
		request: request,
	}
}
