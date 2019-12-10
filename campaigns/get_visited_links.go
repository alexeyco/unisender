package campaigns

import (
	"strconv"

	"github.com/alexeyco/unisender/api"
)

// GetVisitedLinksResult visited links.
type GetVisitedLinksResult struct {
	Fields []string   `json:"fields"`
	Data   [][]string `json:"data"`
}

// GetVisitedLinksRequest request to report on the links visited by users in the specified email campaign.
type GetVisitedLinksRequest struct {
	request *api.Request
}

// Group group results results by the links visited. If the user has visited the link several times, the results
// will be represented by one record indicating the number of visits in the count field.
func (r *GetVisitedLinksRequest) Group() *GetVisitedLinksRequest {
	r.request.Add("group", "1")
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *GetVisitedLinksRequest) Execute() (res *GetVisitedLinksResult, err error) {
	var result GetVisitedLinksResult
	if err = r.request.Execute("getVisitedLinks", &result); err != nil {
		return
	}

	res = &result

	return
}

// GetVisitedLinks returns request to report on the links visited by users in the specified email campaign.
func GetVisitedLinks(request *api.Request, campaignID int64) *GetVisitedLinksRequest {
	request.Add("campaign_id", strconv.FormatInt(campaignID, 10))

	return &GetVisitedLinksRequest{
		request: request,
	}
}
