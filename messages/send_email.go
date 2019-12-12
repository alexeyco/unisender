package messages

import (
	"fmt"
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// SendEmailResult result of sendEmail request.
type SendEmailResult struct {
	EmailID int64 `json:"email_id"`
}

// SendEmailRequest request to send a single individual email without personalization and with limited possibilities
// to obtain statistics.
type SendEmailRequest struct {
	request *api.Request
}

// SenderName sets sender's name.
func (r *SendEmailRequest) SenderName(senderName string) *SendEmailRequest {
	r.request.Add("sender_name", senderName)
	return r
}

// SenderEmail sets sender’s email address. This email must be checked (to do this, you need to manually create
// at least one email with this return address via the web interface, then click on the "send the confirmation request"
// link and follow the link from the email).
func (r *SendEmailRequest) SenderEmail(senderEmail string) *SendEmailRequest {
	r.request.Add("sender_email", senderEmail)
	return r
}

// Subject sets the letter subject.
func (r *SendEmailRequest) Subject(subject string) *SendEmailRequest {
	r.request.Add("subject", subject)
	return r
}

// Body sets text of the letter in the HTML format.
func (r *SendEmailRequest) Body(body string) *SendEmailRequest {
	r.request.Add("body", body)
	return r
}

// ListID code of the list from which the recipient will be offered to unsubscribe in case he follows
// the unsubscribe link.
func (r *SendEmailRequest) ListID(listID int64) *SendEmailRequest {
	r.request.Add("list_id", strconv.FormatInt(listID, 10))
	return r
}

// Attachment
func (r *SendEmailRequest) Attachment(name, content string) *SendEmailRequest {
	r.request.Add(fmt.Sprintf("attachments[%s]", name), content)
	return r
}

// LangDA sets letter language to Danish.
func (r *SendEmailRequest) LangDA() *SendEmailRequest {
	r.request.Add("lang", "da")
	return r
}

// LangDA sets letter language to Danish.
func (r *SendEmailRequest) LangDE() *SendEmailRequest {
	r.request.Add("lang", "de")
	return r
}

// LangDA sets letter language to Spanish.
func (r *SendEmailRequest) LangES() *SendEmailRequest {
	r.request.Add("lang", "es")
	return r
}

// LangDA sets letter language to French.
func (r *SendEmailRequest) LangFR() *SendEmailRequest {
	r.request.Add("lang", "fr")
	return r
}

// LangDA sets letter language to Dutch.
func (r *SendEmailRequest) LangNL() *SendEmailRequest {
	r.request.Add("lang", "nl")
	return r
}

// LangDA sets letter language to Polish.
func (r *SendEmailRequest) LangPL() *SendEmailRequest {
	r.request.Add("lang", "pl")
	return r
}

// LangDA sets letter language to Portuguese.
func (r *SendEmailRequest) LangPT() *SendEmailRequest {
	r.request.Add("lang", "pt")
	return r
}

// LangDA sets letter language to Turkish.
func (r *SendEmailRequest) LangTR() *SendEmailRequest {
	r.request.Add("lang", "tr")
	return r
}

// TrackRead track the fact of reading the email.
func (r *SendEmailRequest) TrackRead() *SendEmailRequest {
	r.request.Add("track_read", "1")
	return r
}

// TrackLinks track whether the recipient has followed the links in the emails.
func (r *SendEmailRequest) TrackLinks() *SendEmailRequest {
	r.request.Add("track_links", "1")
	return r
}

// CC sets an address of secondary recipient of the letter who are sent a copy of the letter. It must contain
// no more than 1 element. You can use the cc parameter to debug or to prove that you have sent the letter.
func (r *SendEmailRequest) CC(cc string) *SendEmailRequest {
	r.request.Add("cc", cc)
	return r
}

// WrapTypeSkip skips letter text alignment.
func (r *SendEmailRequest) WrapTypeSkip() *SendEmailRequest {
	r.request.Add("wrap_type", "skip")
	return r
}

// WrapTypeRight sets letter text alignment to right.
func (r *SendEmailRequest) WrapTypeRight() *SendEmailRequest {
	r.request.Add("wrap_type", "right")
	return r
}

// WrapTypeLeft sets letter text alignment to left.
func (r *SendEmailRequest) WrapTypeLeft() *SendEmailRequest {
	r.request.Add("wrap_type", "left")
	return r
}

// WrapTypeCenter sets letter text alignment to center.
func (r *SendEmailRequest) WrapTypeCenter() *SendEmailRequest {
	r.request.Add("wrap_type", "center")
	return r
}

// ImagesAsAttachments sets the processing mode of the images attached in the letter. In this case images
// will be saved inside the letter as attachments
func (r *SendEmailRequest) ImagesAsAttachments() *SendEmailRequest {
	r.request.Add("images_as", "attachments")
	return r
}

// ImagesAsOnlyLinks sets the processing mode of the images attached in the letter. In this case images sent
// in the image request will be stored on our service, only links to them will be displayed in the letter
// (this will reduce the size of the letter)
func (r *SendEmailRequest) ImagesAsOnlyLinks() *SendEmailRequest {
	r.request.Add("images_as", "only_links")
	return r
}

// ImagesAsUserDefault sets the processing mode of the images attached in the letter. In this case one
// of the above modes will be used which are set for a specific user by the customer support service or by the reseller.
func (r *SendEmailRequest) ImagesAsUserDefault() *SendEmailRequest {
	r.request.Add("images_as", "user_default")
	return r
}

// RefKey sets the parameter can be transferred by the user to assign an identifier key to the letter.
// The obtained key value must be unique.
func (r *SendEmailRequest) RefKey(refKey int64) *SendEmailRequest {
	r.request.Add("ref_key", strconv.FormatInt(refKey, 10))
	return r
}

// MetaData sets metadata sent in the request is returned in webhooks.
func (r *SendEmailRequest) MetaData(name, value string) *SendEmailRequest {
	r.request.Add(fmt.Sprintf("metadata[%s]", name), value)
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *SendEmailRequest) Execute() (emailID int64, err error) {
	var result SendEmailResult
	if err = r.request.Execute("sendEmail", &result); err != nil {
		return
	}

	emailID = result.EmailID

	return
}

// SendEmail returns request to send a single individual email without personalization and with limited possibilities
// to obtain statistics.
func SendEmail(request *api.Request, email string) *SendEmailRequest {
	request.Add("email", email).
		Add("error_checking", "1")

	return &SendEmailRequest{
		request: request,
	}
}
