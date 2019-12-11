package messages

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// DeleteMessageRequest request to delete a message.
type DeleteMessageRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *DeleteMessageRequest) Execute() (err error) {
	err = r.request.Execute("deleteMessage", nil)
	return
}

// DeleteMessage returns request to delete a message.
func DeleteMessage(request *api.Request, messageID int64) *DeleteMessageRequest {
	request.Add("message_id", strconv.FormatInt(messageID, 10))

	return &DeleteMessageRequest{
		request: request,
	}
}
