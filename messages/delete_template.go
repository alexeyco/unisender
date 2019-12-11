package messages

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// DeleteTemplateRequest request to delete a message.
type DeleteTemplateRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *DeleteTemplateRequest) Execute() (err error) {
	err = r.request.Execute("deleteMessage", nil)
	return
}

// DeleteTemplate returns request to delete a template.
func DeleteTemplate(request *api.Request, templateID int64) *DeleteTemplateRequest {
	request.Add("template_id", strconv.FormatInt(templateID, 10))

	return &DeleteTemplateRequest{
		request: request,
	}
}
