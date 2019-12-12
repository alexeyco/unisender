package messages

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// UpdateEmailTemplateRequest request to create an email template for a mass campaign.
type UpdateEmailTemplateRequest struct {
	request *api.Request
}

// Subject sets template title.
func (r *UpdateEmailTemplateRequest) Title(title string) *UpdateEmailTemplateRequest {
	r.request.Add("title", title)
	return r
}

// Subject sets email template subject.
func (r *UpdateEmailTemplateRequest) Subject(subject string) *UpdateEmailTemplateRequest {
	r.request.Add("subject", subject)
	return r
}

// Body sets email body.
func (r *UpdateEmailTemplateRequest) Body(body string) *UpdateEmailTemplateRequest {
	r.request.Add("body", body)
	return r
}

// BodyText sets email body text.
func (r *UpdateEmailTemplateRequest) BodyText(bodyText string) *UpdateEmailTemplateRequest {
	r.request.Add("text_body", bodyText)
	return r
}

// BodyRaw sets email body raw. It is intended to save the json structure of the block editor data structure
// (if the value is message_format = block) The parameter obtains only the JSON structure, otherwise
// it will not be transferred.
func (r *UpdateEmailTemplateRequest) BodyRaw(bodyRaw string) *UpdateEmailTemplateRequest {
	r.request.Add("raw_body", bodyRaw)
	return r
}

// LangDA sets letter language to Danish.
func (r *UpdateEmailTemplateRequest) LangDA() *UpdateEmailTemplateRequest {
	r.request.Add("lang", "da")
	return r
}

// LangDA sets letter language to Danish.
func (r *UpdateEmailTemplateRequest) LangDE() *UpdateEmailTemplateRequest {
	r.request.Add("lang", "de")
	return r
}

// LangDA sets letter language to Spanish.
func (r *UpdateEmailTemplateRequest) LangES() *UpdateEmailTemplateRequest {
	r.request.Add("lang", "es")
	return r
}

// LangDA sets letter language to French.
func (r *UpdateEmailTemplateRequest) LangFR() *UpdateEmailTemplateRequest {
	r.request.Add("lang", "fr")
	return r
}

// LangDA sets letter language to Dutch.
func (r *UpdateEmailTemplateRequest) LangNL() *UpdateEmailTemplateRequest {
	r.request.Add("lang", "nl")
	return r
}

// LangDA sets letter language to Polish.
func (r *UpdateEmailTemplateRequest) LangPL() *UpdateEmailTemplateRequest {
	r.request.Add("lang", "pl")
	return r
}

// LangDA sets letter language to Portuguese.
func (r *UpdateEmailTemplateRequest) LangPT() *UpdateEmailTemplateRequest {
	r.request.Add("lang", "pt")
	return r
}

// LangDA sets letter language to Turkish.
func (r *UpdateEmailTemplateRequest) LangTR() *UpdateEmailTemplateRequest {
	r.request.Add("lang", "tr")
	return r
}

func (r *UpdateEmailTemplateRequest) Description(description string) *UpdateEmailTemplateRequest {
	r.request.Add("description", description)
	return r
}

// MessageFormatBlock sets message format to block.
func (r *UpdateEmailTemplateRequest) MessageFormatBlock() *UpdateEmailTemplateRequest {
	r.request.Add("message_format", "block")
	return r
}

// MessageFormatRawHTML sets message format to raw html.
func (r *UpdateEmailTemplateRequest) MessageFormatRawHTML() *UpdateEmailTemplateRequest {
	r.request.Add("message_format", "raw_html")
	return r
}

// MessageFormatText sets message format to text.
func (r *UpdateEmailTemplateRequest) MessageFormatText() *UpdateEmailTemplateRequest {
	r.request.Add("message_format", "text")
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *UpdateEmailTemplateRequest) Execute() (err error) {
	err = r.request.Execute("updateEmailTemplate", nil)
	return
}

// UpdateEmailTemplate returns request to create an email template for a mass campaign.
func UpdateEmailTemplate(request *api.Request, templateID int64) *UpdateEmailTemplateRequest {
	request.Add("template_id", strconv.FormatInt(templateID, 10))

	return &UpdateEmailTemplateRequest{
		request: request,
	}
}
