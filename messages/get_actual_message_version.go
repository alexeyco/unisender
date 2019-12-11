package messages

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// GetActualMessageVersionResult actual message version.
type GetActualMessageVersionResult struct {
	MessageID       int64 `json:"message_id"`
	ActualVersionID int64 `json:"actual_version_id"`
}

// GetActualMessageVersionRequest request to get actual message version.
type GetActualMessageVersionRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetActualMessageVersionRequest) Execute() (res *GetActualMessageVersionResult, err error) {
	var result GetActualMessageVersionResult
	if err = r.request.Execute("getActualMessageVersion", &result); err != nil {
		return
	}

	res = &result

	return
}

// GetActualMessageVersion returns request to get actual message version.
func GetActualMessageVersion(request *api.Request, messageID int64) *GetActualMessageVersionRequest {
	request.Add("message_id", strconv.FormatInt(messageID, 10))

	return &GetActualMessageVersionRequest{
		request: request,
	}
}
