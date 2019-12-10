package campaigns

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// CancelCampaignRequest request to cancel campaign.
type CancelCampaignRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *CancelCampaignRequest) Execute() (err error) {
	return r.request.Execute("cancelCampaign", nil)
}

// CancelCampaign returns request to cancel a scheduled campaign.
//
// See: https://www.unisender.com/en/support/api/partners/cancel-campaign/
func CancelCampaign(request *api.Request, campaignID int64) *CancelCampaignRequest {
	request.Add("campaign_id", strconv.FormatInt(campaignID, 10))

	return &CancelCampaignRequest{
		request: request,
	}
}
