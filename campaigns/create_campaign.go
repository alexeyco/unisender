package campaigns

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/unisender/api"
)

// CreateCampaignResult result of createCampaign request.
type CreateCampaignResult struct {
	CampaignID int64  `json:"campaign_id"`
	Status     string `json:"status"`
	Count      int64  `json:"count"`
}

// CreateCampaignRequest request to schedule or immediately start sending email or SMS messages.
type CreateCampaignRequest struct {
	request *api.Request
}

// StartTime sets campaign launch date and time. If the argument is not set, the campaign starts immediately.
// To provide additional error protection, you should not schedule two sendings of the same message within an hour.
func (r *CreateCampaignRequest) StartTime(startTime time.Time) *CreateCampaignRequest {
	r.request.Add("start_time", startTime.UTC().Format("2006-01-02 15:04")).
		Add("time_zone", "UTC")

	return r
}

// TrackRead enables tracking the fact of reading the email message.
func (r *CreateCampaignRequest) TrackRead() *CreateCampaignRequest {
	r.request.Add("track_read", "1")
	return r
}

// TrackLinks enables tracking click-throughs in email messages.
func (r *CreateCampaignRequest) TrackLinks() *CreateCampaignRequest {
	r.request.Add("track_links", "1")
	return r
}

// Contacts sets email addresses to which sending of a message should be limited.
func (r *CreateCampaignRequest) Contacts(emails ...string) *CreateCampaignRequest {
	r.request.Add("contacts", strings.Join(emails, ","))
	return r
}

// ContactsURL instead of the contacts parameter containing the actual email addresses or phone numbers,
// in this parameter you can specify the URL of the file from which the addresses (phone numbers) will be read.
//
// The URL must start with “http://”, “https://” or “ftp://”. The file must contain one contact per string,
// without commas; strings must be separated by “n” or “rn” (Mac format — only “r” — not supported).
// The file can be deleted after the campaign has shifted to the ‘scheduled’ status.
func (r *CreateCampaignRequest) ContactsURL(contactsURL string) *CreateCampaignRequest {
	r.request.Add("contacts_url", contactsURL)
	return r
}

// TrackGoogleAnalytics enables Google Analytics/Yandex.Metrica integration for this campaign.
func (r *CreateCampaignRequest) TrackGoogleAnalytics() *CreateCampaignRequest {
	r.request.Add("track_ga", "1")
	return r
}

// GoogleAnalyticsMedium sets Google Analytics medium parameter. Medium is broad bucket of categories that describe
// the kind of traffic being driven to.
func (r *CreateCampaignRequest) GoogleAnalyticsMedium(medium string) *CreateCampaignRequest {
	r.request.Add("ga_medium", medium)
	return r
}

// GoogleAnalyticsSource sets Google Analytics source parameter. Source is the actual domain sending traffic to.
func (r *CreateCampaignRequest) GoogleAnalyticsSource(source string) *CreateCampaignRequest {
	r.request.Add("ga_source", source)
	return r
}

// GoogleAnalyticsCampaign sets Google Analytics source parameter.
func (r *CreateCampaignRequest) GoogleAnalyticsCampaign(campaign string) *CreateCampaignRequest {
	r.request.Add("ga_campaign", campaign)
	return r
}

// GoogleAnalyticsContent sets Google Analytics content parameter.
func (r *CreateCampaignRequest) GoogleAnalyticsContent(content string) *CreateCampaignRequest {
	r.request.Add("ga_content", content)
	return r
}

// GoogleAnalyticsTerm sets Google Analytics term parameter.
func (r *CreateCampaignRequest) GoogleAnalyticsTerm(term string) *CreateCampaignRequest {
	r.request.Add("ga_term", term)
	return r
}

// Payment sets payment limit and currency. Allows you to limit the campaign budget.
func (r *CreateCampaignRequest) Payment(limit float64, currency string) *CreateCampaignRequest {
	r.request.Add("payment_limit", fmt.Sprintf("%f", limit)).
		Add("payment_currency", currency)

	return r
}

// Execute sends request to UniSender API and returns result.
func (r *CreateCampaignRequest) Execute() (res *CreateCampaignResult, err error) {
	var result CreateCampaignResult
	if err = r.request.Execute("createCampaign", &result); err != nil {
		return
	}

	res = &result

	return
}

// CreateCampaign returns request to schedule or immediately start sending email or SMS messages.
func CreateCampaign(request *api.Request, messageID int64) *CreateCampaignRequest {
	request.Add("message_id", strconv.FormatInt(messageID, 10))

	return &CreateCampaignRequest{
		request: request,
	}
}
