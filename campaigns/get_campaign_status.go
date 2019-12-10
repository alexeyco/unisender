package campaigns

import (
	"strconv"
	"time"

	"github.com/alexeyco/unisender/api"
)

// GetCampaignStatusResult campaign status result.
type GetCampaignStatusResult struct {
	Status       string    `json:"status"`
	CreationTime time.Time `json:"creation_time"`
	StartTime    time.Time `json:"start_time"`
}

// GetCampaignStatusRequest request to find out the status of the campaign.
type GetCampaignStatusRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetCampaignStatusRequest) Execute() (res *GetCampaignStatusResult, err error) {
	var result GetCampaignStatusResult
	if err = r.request.Execute("getCampaignStatus", &result); err != nil {
		return
	}

	res = &result

	return
}

// GetCampaignStatus returns request to find out the status of the campaign.
func GetCampaignStatus(request *api.Request, campaignID int64) *GetCampaignStatusRequest {
	request.Add("campaign_id", strconv.FormatInt(campaignID, 10))

	return &GetCampaignStatusRequest{
		request: request,
	}
}
