package campaigns

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// GetCampaignStatsResult campaign statistics.
type GetCampaignStatsResult struct {
	Total         int64 `json:"total"`
	Sent          int64 `json:"sent"`
	Delivered     int64 `json:"delivered"`
	ReadUnique    int64 `json:"read_unique"`
	ReadAll       int64 `json:"read_all"`
	ClickedUnique int64 `json:"clicked_unique"`
	ClickedAll    int64 `json:"clicked_all"`
	Unsubscribed  int64 `json:"unsubscribed"`
	Spam          int64 `json:"spam"`
}

// GetCampaignCommonStatsRequest
type GetCampaignCommonStatsRequest struct {
	request *api.Request
}

// Execute sends request to UniSender API and returns result.
func (r *GetCampaignCommonStatsRequest) Execute() (res *GetCampaignStatsResult, err error) {
	var result GetCampaignStatsResult
	if err = r.request.Execute("getCampaignCommonStats", &result); err != nil {
		return
	}

	res = &result

	return
}

// GetCampaignCommonStats general information about the results of delivering messages in the given list.
func GetCampaignCommonStats(request *api.Request, campaignID int64) *GetCampaignCommonStatsRequest {
	request.Add("campaign_id", strconv.FormatInt(campaignID, 10))

	return &GetCampaignCommonStatsRequest{
		request: request,
	}
}
