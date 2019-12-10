package unisender

import (
	"github.com/alexeyco/unisender/contacts"
	"net/http"
	"sync"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/campaigns"
	"github.com/alexeyco/unisender/common"
	"github.com/alexeyco/unisender/contacts2"
	"github.com/alexeyco/unisender/messages2"
)

// LanguageDefault default API response language.
const LanguageDefault = "en"

// UniSender API client struct.
type UniSender struct {
	apiKey   string
	language string
	client   *http.Client
	logger   api.Logger
	mu       sync.RWMutex
}

// SetLanguageEnglish sets API response language to English.
func (u *UniSender) SetLanguageEnglish() *UniSender {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.language = "en"

	return u
}

// SetLanguageItalian sets API response language to Italian.
func (u *UniSender) SetLanguageItalian() *UniSender {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.language = "it"

	return u
}

// SetLanguageRussian sets API response language to Russian.
func (u *UniSender) SetLanguageRussian() *UniSender {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.language = "ru"

	return u
}

// SetClient sets custom http.Request to UniSender client.
func (u *UniSender) SetClient(client *http.Client) *UniSender {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.client = client

	return u
}

// SetLogger sets logger to UniSender client.
func (u *UniSender) SetLogger(logger api.Logger) *UniSender {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.logger = logger

	return u
}

// CancelCampaign returns request to cancel a scheduled campaign.
//
// See: https://www.unisender.com/en/support/api/partners/cancel-campaign/
//
// A campaign can be canceled if it has one of the following statuses:
//   * "Planned",
//   * "Scheduled" (scheduled),
//   * "Considered by the administration" (censor_hold),
//   * "Waiting for approval" (waits_censor).
//
// The distribution status can be obtained using the getCampaignStatus method.
//
// See: https://www.unisender.com/en/support/api/partners/getcampaignstatus/
func (u *UniSender) CancelCampaign(campaignID int64) *campaigns.CancelCampaignRequest {
	return campaigns.CancelCampaign(u.request(), campaignID)
}

// CreateCampaign returns request used to schedule or immediately start sending email or SMS messages. The same message
// can be sent several times, but the moments of sending should have an interval of at least an hour.
//
// See: https://www.unisender.com/en/support/api/partners/createcampaign/
//
// This method is used to send the already created messages. To create messages in advance, use the createEmailMessage
// and createSmsMessage methods.
//
// See: https://www.unisender.com/en/support/api/partners/createemailmessage/
// See: https://www.unisender.com/en/support/api/partners/createsmsmessage/
func (u *UniSender) CreateCampaign(messageID int64) *campaigns.CreateCampaignRequest {
	return campaigns.CreateCampaign(u.request(), messageID)
}

// GetCampaignCommonStats general information about the results of delivering messages in the given list.
//
// See: https://www.unisender.com/en/support/api/partners/get-campaign-common-stats/
func (u *UniSender) GetCampaignCommonStats(campaignID int64) *campaigns.GetCampaignCommonStatsRequest {
	return campaigns.GetCampaignCommonStats(u.request(), campaignID)
}

// GetCampaigns returns request to get the list of all available campaigns. The number of mailings received at a time
// is limited to 10,000. To get a complete mailings list when there is more than 10,000, use the from and to parameters.
//
// See: https://www.unisender.com/en/support/statistics/getcampaigns/
func (u *UniSender) GetCampaigns() *campaigns.GetCampaignsRequest {
	return campaigns.GetCampaigns(u.request())
}

// GetCampaignStatus returns request to find out the status of the campaign.
//
// Campaign status possible options:
//   waits_censor   — campaign is waiting to be checked.
//
//   censor_hold    — it is actually equivalent to waits_censor: considered by the administrator, but delayed
//                    for further check.
//
//   declined       — the campaign has been rejected by administrator.
//
//   waits_schedule — the task for placing the list in the queue has been received and the campaign is waiting
//                    to be placed in the queue. As a rule, the campaign stays in this status one or two minutes
//                    before changing its status on scheduled.
//
//   scheduled      — scheduled to be launched. It will be launched as soon as the sending time comes.
//
//   in_progress    — messages are being sent.
//
//   analysed       — all messages have been sent, the results are being analyzed.
//
//   completed      — all messages have been sent and analysis of the results is completed.
//
//   stopped        — the campaign is paused.
//
//   canceled       — the campaign is canceled (usually due to the lack of money or at the request of the user).
//
// See: https://www.unisender.com/en/support/api/partners/getcampaignstatus/
func (u *UniSender) GetCampaignStatus(campaignID int64) *campaigns.GetCampaignStatusRequest {
	return campaigns.GetCampaignStatus(u.request(), campaignID)
}

// GetVisitedLinks returns request to report on the links visited by users in the specified email campaign.
//
// See: https://www.unisender.com/en/support/api/partners/getvisitedlinks/
func (u *UniSender) GetVisitedLinks(campaignID int64) *campaigns.GetVisitedLinksRequest {
	return campaigns.GetVisitedLinks(u.request(), campaignID)
}

// GetCurrencyRates allows you to get a list of all currencies in the UniSender system.
//
// See: https://www.unisender.com/en/support/api/common/getcurrencyrates/
func (u *UniSender) GetCurrencyRates() *common.GetCurrencyRatesRequest {
	return common.GetCurrencyRates(u.request())
}

// CheckEmail returns request to check the delivery status of emails sent using the sendEmail method.
//
// To speed up the work of the sendEmail method, delivery statuses are stored for a limited period of time,
// i.e. only for a month.
//
// See: https://www.unisender.com/en/support/api/messages/checkemail/
func (u *UniSender) CheckEmail(emailIDs ...int64) *contacts.CheckEmailRequest {
	return contacts.CheckEmail(u.request(), emailIDs...)
}

// CreateList creates a new contact list.
func (u *UniSender) CreateList(title string) *contacts2.CreateListRequest {
	return contacts2.CreateList(u.request(), title)
}

// GetLists returns all available campaign lists.
func (u *UniSender) GetLists() *contacts2.GetListsRequest {
	return contacts2.GetLists(u.request())
}

// UpdateList changes campaign list properties.
func (u *UniSender) UpdateList(listID int64, title string) *contacts2.UpdateListRequest {
	return contacts2.UpdateList(u.request(), listID, title)
}

// DeleteList removes a list.
func (u *UniSender) DeleteList(listID int64) *contacts2.DeleteListRequest {
	return contacts2.DeleteList(u.request(), listID)
}

// GetContact returns information about single contact.
func (u *UniSender) GetContact(email string) *contacts2.GetContactRequest {
	return contacts2.GetContact(u.request(), email)
}

// Subscribe subscribes the contact email or phone number to one or several lists.
func (u *UniSender) Subscribe() *contacts2.SubscribeRequest {
	return contacts2.Subscribe(u.request())
}

// Unsubscribe unsubscribes the contact email or phone number from one or several lists.
func (u *UniSender) Unsubscribe(contact string) *contacts2.UnsubscribeRequest {
	return contacts2.Unsubscribe(u.request(), contact)
}

// Exclude excludes the contact’s email or phone number from one or several lists.
func (u *UniSender) Exclude(contact string) *contacts2.ExcludeRequest {
	return contacts2.Exclude(u.request(), contact)
}

// ImportContacts imports contacts.
func (u *UniSender) ImportContacts(collection *contacts2.ImportContactsCollection) *contacts2.ImportContactsRequest {
	return contacts2.ImportContacts(u.request(), collection)
}

// ExportContacts export of contact data from UniSender.
func (u *UniSender) ExportContacts() *contacts2.ExportContactsRequest {
	return contacts2.ExportContacts(u.request())
}

// IsContactInList checks whether the contact is in the specified user lists.
func (u *UniSender) IsContactInList(email string, listIDs ...int64) *contacts2.IsContactInListRequest {
	return contacts2.IsContactInList(u.request(), email, listIDs...)
}

// GetContactCount returns the contacts list size.
func (u *UniSender) GetContactCount(listID int64) *contacts2.GetContactCountRequest {
	return contacts2.GetContactCount(u.request(), listID)
}

// GetTotalContactsCount returns the contacts database size by the user login.
func (u *UniSender) GetTotalContactsCount(login string) *contacts2.GetTotalContactsCountRequest {
	return contacts2.GetTotalContactsCount(u.request(), login)
}

// GetCheckedEmail returns request to check the delivery status of emails sent using.
func (u *UniSender) GetCheckedEmail(login string) *messages2.GetCheckedEmailRequest {
	return messages2.GetCheckedEmail(u.request(), login)
}

// ValidateSender returns request that sends a message to the email address with a link to confirm the address
// as the return address.
func (u *UniSender) ValidateSender(email string) *messages2.ValidateSenderRequest {
	return messages2.ValidateSender(u.request(), email)
}

// GetSenderDomainList returns information about domains.
func (u *UniSender) GetSenderDomainList(login string) *messages2.GetSenderDomainListRequest {
	return messages2.GetSenderDomainList(u.request(), login)
}

// SetSenderDomain register the domain in the list.
func (u *UniSender) SetSenderDomain(login, domain string) *messages2.SetSenderDomainRequest {
	return messages2.SetSenderDomain(u.request(), login, domain)
}

func (u *UniSender) request() *api.Request {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return api.NewRequest(u.client, u.language).
		SetLogger(u.logger).
		Add("format", "json").
		Add("lang", u.language).
		Add("api_key", u.apiKey)
}

// New returns new UniSender API client.
func New(apiKey string) *UniSender {
	return &UniSender{
		apiKey:   apiKey,
		language: LanguageDefault,
		client:   http.DefaultClient,
		mu:       sync.RWMutex{},
	}
}
