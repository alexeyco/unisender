package messages

import (
	"strconv"
	"time"

	"github.com/alexeyco/unisender/api"
)

// GetTemplateResult information about SMS or email message.
type GetTemplateResult struct {
	ID                    int64       `json:"id"`
	SubUserLogin          string      `json:"sub_user_login"`
	Title                 string      `json:"title"`
	Description           string      `json:"description"`
	Lang                  string      `json:"lang_code"`
	Subject               string      `json:"subject"`
	Attachments           interface{} `json:"attachments"`
	ScreenshotURL         string      `json:"screenshot_url"`
	FullsizeScreenshotURL string      `json:"fullsize_screenshot_url"`
	Created               time.Time   `json:"created"`
	MessageFormat         string      `json:"message_format"`
	Type                  string      `json:"type"`
	Body                  string      `json:"body"`
	BodyRaw               string      `json:"raw_body"`
}

// GetTemplateRequest request to get information about SMS or email message.
type GetTemplateRequest struct {
	request *api.Request
}

// SystemTemplateID sets system template ID. It is returned by the getTemplates or listTemplates methods.
func (r *GetTemplateRequest) SystemTemplateID(systemTemplateID int64) *GetTemplateRequest {
	r.request.Add("system_template_id", strconv.FormatInt(systemTemplateID, 10))
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *GetTemplateRequest) Execute() (res *GetTemplateResult, err error) {
	var result GetTemplateResult
	if err = r.request.Execute("getTemplate", &result); err != nil {
		return
	}

	res = &result

	return
}

// GetTemplate returns request to get information about the specified template.
func GetTemplate(request *api.Request, messageID int64) *GetTemplateRequest {
	request.Add("template_id", strconv.FormatInt(messageID, 10))

	return &GetTemplateRequest{
		request: request,
	}
}
