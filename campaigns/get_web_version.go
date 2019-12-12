package campaigns

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

type GetWebVersionResult struct {
	LetterID      int64  `json:"letter_id"`
	WebLetterLink string `json:"web_letter_link"`
}

type GetWebVersionRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetWebVersionRequest) Execute() (res *GetWebVersionResult, err error) {
	var result GetWebVersionResult
	if err = r.request.Execute("getWebVersion", &result); err != nil {
		return
	}

	res = &result

	return
}

// GetWebVersion returns request to get the link to the web version of the letter.
func GetWebVersion(request *api.Request, campaignID int64) *GetWebVersionRequest {
	request.Add("campaign_id", strconv.FormatInt(campaignID, 10))

	return &GetWebVersionRequest{
		request: request,
	}
}
