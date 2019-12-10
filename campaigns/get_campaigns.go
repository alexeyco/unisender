package campaigns

import (
	"strconv"
	"time"

	"github.com/alexeyco/unisender/api"
)

// GetCampaignsResult campaign info.
type GetCampaignsResult struct {
	ID          int64     `json:"id"`
	StartTime   time.Time `json:"start_time"`
	Status      string    `json:"status"`
	MessageID   int64     `json:"message_id"`
	ListID      int64     `json:"list_id"`
	Subject     string    `json:"subject"`
	SenderName  string    `json:"sender_name"`
	SenderEmail string    `json:"sender_email"`
	StatsURL    string    `json:"stats_url"`
}

// GetCampaignsRequest request to get the list of all available campaigns.
type GetCampaignsRequest struct {
	request *api.Request
}

// From sets time of the campaign start from which campaigns are to be displayed.
func (r *GetCampaignsRequest) From(from time.Time) *GetCampaignsRequest {
	r.request.Add("from", from.UTC().Format("2006-01-02 15:04:05"))
	return r
}

// To sets time of the campaign start until which campaigns are to be displayed.
func (r *GetCampaignsRequest) To(to time.Time) *GetCampaignsRequest {
	r.request.Add("to", to.UTC().Format("2006-01-02 15:04:05"))
	return r
}

// Limit sets number of records in the response.
func (r *GetCampaignsRequest) Limit(limit int) *GetCampaignsRequest {
	r.request.Add("limit", strconv.Itoa(limit))
	return r
}

// Offset sets offset from which position the selection is to be started.
func (r *GetCampaignsRequest) Offset(offset int) *GetCampaignsRequest {
	r.request.Add("offset", strconv.Itoa(offset))
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *GetCampaignsRequest) Execute() (res []GetCampaignsResult, err error) {
	err = r.request.Execute("getCampaigns", &res)
	return
}

// GetCampaigns returns request to get the list of all available campaigns.
func GetCampaigns(request *api.Request) *GetCampaignsRequest {
	return &GetCampaignsRequest{
		request: request,
	}
}
