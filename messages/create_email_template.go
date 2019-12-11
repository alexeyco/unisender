package messages

import "github.com/alexeyco/unisender/api"

// CreateEmailTemplateResult result of createEmailTemplate request.
type CreateEmailTemplateResult struct {
	TemplateID int64 `json:"template_id"`
}

// CreateEmailTemplateRequest request to create an email template for a mass campaign.
type CreateEmailTemplateRequest struct {
	request *api.Request
}

// Subject sets email subject.
func (r *CreateEmailTemplateRequest) Subject(subject string) *CreateEmailTemplateRequest {
	r.request.Add("subject", subject)
	return r
}

// Body sets email body.
func (r *CreateEmailTemplateRequest) Body(body string) *CreateEmailTemplateRequest {
	r.request.Add("body", body)
	return r
}

// BodyText sets email body text.
func (r *CreateEmailTemplateRequest) BodyText(bodyText string) *CreateEmailTemplateRequest {
	r.request.Add("text_body", bodyText)
	return r
}

// BodyRaw sets email body raw. It is intended to save the json structure of the block editor data structure
// (if the value is message_format = block) The parameter obtains only the JSON structure, otherwise
// it will not be transferred.
func (r *CreateEmailTemplateRequest) BodyRaw(bodyRaw string) *CreateEmailTemplateRequest {
	r.request.Add("raw_body", bodyRaw)
	return r
}

// LangDA sets letter language to Danish.
func (r *CreateEmailTemplateRequest) LangDA() *CreateEmailTemplateRequest {
	r.request.Add("lang", "da")
	return r
}

// LangDA sets letter language to Danish.
func (r *CreateEmailTemplateRequest) LangDE() *CreateEmailTemplateRequest {
	r.request.Add("lang", "de")
	return r
}

// LangDA sets letter language to Spanish.
func (r *CreateEmailTemplateRequest) LangES() *CreateEmailTemplateRequest {
	r.request.Add("lang", "es")
	return r
}

// LangDA sets letter language to French.
func (r *CreateEmailTemplateRequest) LangFR() *CreateEmailTemplateRequest {
	r.request.Add("lang", "fr")
	return r
}

// LangDA sets letter language to Dutch.
func (r *CreateEmailTemplateRequest) LangNL() *CreateEmailTemplateRequest {
	r.request.Add("lang", "nl")
	return r
}

// LangDA sets letter language to Polish.
func (r *CreateEmailTemplateRequest) LangPL() *CreateEmailTemplateRequest {
	r.request.Add("lang", "pl")
	return r
}

// LangDA sets letter language to Portuguese.
func (r *CreateEmailTemplateRequest) LangPT() *CreateEmailTemplateRequest {
	r.request.Add("lang", "pt")
	return r
}

// LangDA sets letter language to Turkish.
func (r *CreateEmailTemplateRequest) LangTR() *CreateEmailTemplateRequest {
	r.request.Add("lang", "tr")
	return r
}

func (r *CreateEmailTemplateRequest) Description(description string) *CreateEmailTemplateRequest {
	r.request.Add("description", description)
	return r
}

// MessageFormatBlock sets message format to block.
func (r *CreateEmailTemplateRequest) MessageFormatBlock() *CreateEmailTemplateRequest {
	r.request.Add("message_format", "block")
	return r
}

// MessageFormatRawHTML sets message format to raw html.
func (r *CreateEmailTemplateRequest) MessageFormatRawHTML() *CreateEmailTemplateRequest {
	r.request.Add("message_format", "raw_html")
	return r
}

// MessageFormatText sets message format to text.
func (r *CreateEmailTemplateRequest) MessageFormatText() *CreateEmailTemplateRequest {
	r.request.Add("message_format", "text")
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *CreateEmailTemplateRequest) Execute() (templateID int64, err error) {
	var result CreateEmailTemplateResult
	if err = r.request.Execute("createEmailTemplate", &result); err != nil {
		return
	}

	templateID = result.TemplateID

	return
}

// CreateEmailTemplate returns request to create an email template for a mass campaign.
func CreateEmailTemplate(request *api.Request, title string) *CreateEmailTemplateRequest {
	request.Add("title", title)

	return &CreateEmailTemplateRequest{
		request: request,
	}
}
