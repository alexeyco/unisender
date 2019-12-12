package messages

import (
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

type UpdateEmailMessageRequest struct {
	request *api.Request
}

// SenderName sets sender name.
func (r *UpdateEmailMessageRequest) SenderName(senderName string) *UpdateEmailMessageRequest {
	r.request.Add("sender_name", senderName)
	return r
}

// SenderEmail sets sender email.
func (r *UpdateEmailMessageRequest) SenderEmail(senderEmail string) *UpdateEmailMessageRequest {
	r.request.Add("sender_email", senderEmail)
	return r
}

// Subject sets email subject.
func (r *UpdateEmailMessageRequest) Subject(subject string) *UpdateEmailMessageRequest {
	r.request.Add("subject", subject)
	return r
}

// Body sets email body.
func (r *UpdateEmailMessageRequest) Body(body string) *UpdateEmailMessageRequest {
	r.request.Add("body", body)
	return r
}

// BodyText sets email text body.
func (r *UpdateEmailMessageRequest) BodyText(bodyText string) *UpdateEmailMessageRequest {
	r.request.Add("text_body", bodyText)
	return r
}

// BodyRaw sets email raw body.
func (r *UpdateEmailMessageRequest) BodyRaw(bodyRaw string) *UpdateEmailMessageRequest {
	r.request.Add("raw_body", bodyRaw)
	return r
}

// ListID sets list ID.
func (r *UpdateEmailMessageRequest) ListID(listID int64) *UpdateEmailMessageRequest {
	r.request.Add("list_id", strconv.FormatInt(listID, 10))
	return r
}

// MessageFormatBlock sets message format to block.
func (r *UpdateEmailMessageRequest) MessageFormatBlock() *UpdateEmailMessageRequest {
	r.request.Add("message_format", "block")
	return r
}

// MessageFormatRawHTML sets message format to raw html.
func (r *UpdateEmailMessageRequest) MessageFormatRawHTML() *UpdateEmailMessageRequest {
	r.request.Add("message_format", "raw_html")
	return r
}

// MessageFormatText sets message format to text.
func (r *UpdateEmailMessageRequest) MessageFormatText() *UpdateEmailMessageRequest {
	r.request.Add("message_format", "text")
	return r
}

// LangDA sets letter language to Danish.
func (r *UpdateEmailMessageRequest) LangDA() *UpdateEmailMessageRequest {
	r.request.Add("lang", "da")
	return r
}

// LangDA sets letter language to Danish.
func (r *UpdateEmailMessageRequest) LangDE() *UpdateEmailMessageRequest {
	r.request.Add("lang", "de")
	return r
}

// LangDA sets letter language to Spanish.
func (r *UpdateEmailMessageRequest) LangES() *UpdateEmailMessageRequest {
	r.request.Add("lang", "es")
	return r
}

// LangDA sets letter language to French.
func (r *UpdateEmailMessageRequest) LangFR() *UpdateEmailMessageRequest {
	r.request.Add("lang", "fr")
	return r
}

// LangDA sets letter language to Dutch.
func (r *UpdateEmailMessageRequest) LangNL() *UpdateEmailMessageRequest {
	r.request.Add("lang", "nl")
	return r
}

// LangDA sets letter language to Polish.
func (r *UpdateEmailMessageRequest) LangPL() *UpdateEmailMessageRequest {
	r.request.Add("lang", "pl")
	return r
}

// LangDA sets letter language to Portuguese.
func (r *UpdateEmailMessageRequest) LangPT() *UpdateEmailMessageRequest {
	r.request.Add("lang", "pt")
	return r
}

// LangDA sets letter language to Turkish.
func (r *UpdateEmailMessageRequest) LangTR() *UpdateEmailMessageRequest {
	r.request.Add("lang", "tr")
	return r
}

// Categories sets message categories.
func (r *UpdateEmailMessageRequest) Categories(categories ...string) *UpdateEmailMessageRequest {
	r.request.Add("categories", strings.Join(categories, ","))
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *UpdateEmailMessageRequest) Execute() (err error) {
	err = r.request.Execute("updateEmailMessage", nil)
	return
}

// UpdateEmailMessage returns request to change existing email message.
func UpdateEmailMessage(request *api.Request, messageID int64) *UpdateEmailMessageRequest {
	request.Add("id", strconv.FormatInt(messageID, 10))

	return &UpdateEmailMessageRequest{
		request: request,
	}
}
