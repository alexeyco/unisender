package messages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexeyco/unisender/api"
)

// CreateEmailMessageResult createMessage request result.
type CreateEmailMessageResult struct {
	MessageID int64 `json:"message_id"`
}

// CreateEmailMessageRequest request o create an email without sending it.
type CreateEmailMessageRequest struct {
	request *api.Request
}

// SenderName sets sender name.
func (r *CreateEmailMessageRequest) SenderName(senderName string) *CreateEmailMessageRequest {
	r.request.Add("sender_name", senderName)
	return r
}

// SenderEmail sets sender email.
func (r *CreateEmailMessageRequest) SenderEmail(senderEmail string) *CreateEmailMessageRequest {
	r.request.Add("sender_email", senderEmail)
	return r
}

// Subject sets email subject.
func (r *CreateEmailMessageRequest) Subject(subject string) *CreateEmailMessageRequest {
	r.request.Add("subject", subject)
	return r
}

// Body sets email body.
func (r *CreateEmailMessageRequest) Body(body string) *CreateEmailMessageRequest {
	r.request.Add("body", body)
	return r
}

// BodyText sets email body text.
func (r *CreateEmailMessageRequest) BodyText(bodyText string) *CreateEmailMessageRequest {
	r.request.Add("text_body", bodyText)
	return r
}

// BodyRaw sets email body raw. It is intended to save the json structure of the block editor data structure
// (if the value is message_format = block) The parameter obtains only the JSON structure, otherwise
// it will not be transferred.
func (r *CreateEmailMessageRequest) BodyRaw(bodyRaw string) *CreateEmailMessageRequest {
	r.request.Add("raw_body", bodyRaw)
	return r
}

// GenerateText if used, the text part of the letter will be generated automatically based on the HTML part.
// If you do not provide the text version along with the HTML version, you are recommended to use it
// for automatic generation of the text part of the letter.
//
// If the text variant of the letter is provided using the text_body parameter, the generate_text parameter is ignored.
// Thus, if the generate_text value has been set to 1, the server’s response will contain a warning.
func (r *CreateEmailMessageRequest) GenerateText() *CreateEmailMessageRequest {
	r.request.Add("generate_text", "1")
	return r
}

// MessageFormatBlock sets message format to block.
func (r *CreateEmailMessageRequest) MessageFormatBlock() *CreateEmailMessageRequest {
	r.request.Add("message_format", "block")
	return r
}

// MessageFormatRawHTML sets message format to raw html.
func (r *CreateEmailMessageRequest) MessageFormatRawHTML() *CreateEmailMessageRequest {
	r.request.Add("message_format", "raw_html")
	return r
}

// MessageFormatText sets message format to text.
func (r *CreateEmailMessageRequest) MessageFormatText() *CreateEmailMessageRequest {
	r.request.Add("message_format", "text")
	return r
}

// Tag sets message tag. If used, the letter will not be sent on the entire list, but only to those recipients
// to whom the tag is assigned.
func (r *CreateEmailMessageRequest) Tag(tag string) *CreateEmailMessageRequest {
	r.request.Add("tag", tag)
	return r
}

// Attachment attaches file to letter.
func (r *CreateEmailMessageRequest) Attachment(name, content string) *CreateEmailMessageRequest {
	r.request.Add(fmt.Sprintf("attachments[%s]", name), content)
	return r
}

// LangDA sets letter language to Danish.
func (r *CreateEmailMessageRequest) LangDA() *CreateEmailMessageRequest {
	r.request.Add("lang", "da")
	return r
}

// LangDA sets letter language to Danish.
func (r *CreateEmailMessageRequest) LangDE() *CreateEmailMessageRequest {
	r.request.Add("lang", "de")
	return r
}

// LangDA sets letter language to Spanish.
func (r *CreateEmailMessageRequest) LangES() *CreateEmailMessageRequest {
	r.request.Add("lang", "es")
	return r
}

// LangDA sets letter language to French.
func (r *CreateEmailMessageRequest) LangFR() *CreateEmailMessageRequest {
	r.request.Add("lang", "fr")
	return r
}

// LangDA sets letter language to Dutch.
func (r *CreateEmailMessageRequest) LangNL() *CreateEmailMessageRequest {
	r.request.Add("lang", "nl")
	return r
}

// LangDA sets letter language to Polish.
func (r *CreateEmailMessageRequest) LangPL() *CreateEmailMessageRequest {
	r.request.Add("lang", "pl")
	return r
}

// LangDA sets letter language to Portuguese.
func (r *CreateEmailMessageRequest) LangPT() *CreateEmailMessageRequest {
	r.request.Add("lang", "pt")
	return r
}

// LangDA sets letter language to Turkish.
func (r *CreateEmailMessageRequest) LangTR() *CreateEmailMessageRequest {
	r.request.Add("lang", "tr")
	return r
}

// TemplateID sets ID of the user letter template created before
func (r *CreateEmailMessageRequest) TemplateID(templateID int64) *CreateEmailMessageRequest {
	r.request.Add("template_id", strconv.FormatInt(templateID, 10))
	return r
}

// SystemTemplateID sets ID of the user letter system template created before.
func (r *CreateEmailMessageRequest) SystemTemplateID(systemTemplateID int64) *CreateEmailMessageRequest {
	r.request.Add("system_template_id", strconv.FormatInt(systemTemplateID, 10))
	return r
}

// WrapTypeSkip skips letter text alignment.
func (r *CreateEmailMessageRequest) WrapTypeSkip() *CreateEmailMessageRequest {
	r.request.Add("wrap_type", "skip")
	return r
}

// WrapTypeRight sets letter text alignment to right.
func (r *CreateEmailMessageRequest) WrapTypeRight() *CreateEmailMessageRequest {
	r.request.Add("wrap_type", "right")
	return r
}

// WrapTypeLeft sets letter text alignment to left.
func (r *CreateEmailMessageRequest) WrapTypeLeft() *CreateEmailMessageRequest {
	r.request.Add("wrap_type", "left")
	return r
}

// WrapTypeCenter sets letter text alignment to center.
func (r *CreateEmailMessageRequest) WrapTypeCenter() *CreateEmailMessageRequest {
	r.request.Add("wrap_type", "center")
	return r
}

// Categories sets message categories.
func (r *CreateEmailMessageRequest) Categories(categories ...string) *CreateEmailMessageRequest {
	r.request.Add("categories", strings.Join(categories, ","))
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *CreateEmailMessageRequest) Execute() (res int64, err error) {
	var result CreateEmailMessageResult
	if err = r.request.Execute("createEmailMessage", &result); err != nil {
		return
	}

	res = result.MessageID

	return
}

// CreateEmailMessage returns request o create an email without sending it.
func CreateEmailMessage(request *api.Request, listID int64) *CreateEmailMessageRequest {
	request.Add("list_id", strconv.FormatInt(listID, 10))

	return &CreateEmailMessageRequest{
		request: request,
	}
}
