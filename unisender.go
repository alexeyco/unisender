package unisender

import (
	"github.com/alexeyco/unisender/contacts"
	"net/http"
	"sync"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/campaigns"
	"github.com/alexeyco/unisender/common"
	"github.com/alexeyco/unisender/contacts2"
	"github.com/alexeyco/unisender/messages"
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

// CreateField returns request to create a new user field, the value of which can be set for each recipient,
// and then it can be substituted in the letter.
//
// See: https://www.unisender.com/en/support/api/partners/createfield/
func (u *UniSender) CreateField(name string) *contacts.CreateFieldRequest {
	return contacts.CreateField(u.request(), name)
}

// DeleteField returns request to delete a user field.
//
// See: https://www.unisender.com/en/support/api/partners/deletefield/
func (u *UniSender) DeleteField(fieldID int64) *contacts.DeleteFieldRequest {
	return contacts.DeleteField(u.request(), fieldID)
}

// DeleteTag returns request to delete a user tag.
//
// See: https://www.unisender.com/en/support/api/partners/deletetag/
func (u *UniSender) DeleteTag(fieldID int64) *contacts.DeleteTagRequest {
	return contacts.DeleteTag(u.request(), fieldID)
}

// CheckEmail returns request to check the delivery status of emails sent using the sendEmail method.
//
// To speed up the work of the sendEmail method, delivery statuses are stored for a limited period of time,
// i.e. only for a month.
//
// Status possible options:
//   not_sent                      – The message has not been processed yet.
//
//   ok_sent                       – The message has been sent, it is in the intermediate status until the receipt
//                                   of the delivery/non-delivery response.
//
//   ok_delivered                  – The message has been delivered. It may change to ‘ok_read’, ‘ok_link_visited’,
//                                   ‘ok_unsubscribed’ or ‘ok_spam_folder’.
//
//   ok_read                       – The message has been delivered and its reading has been registered. It may change
//                                   to ‘ok_link_visited’, ‘ok_unsubscribed’ or ‘ok_spam_folder’.
//
//   ok_fbl                        – The message has been delivered but placed in the spam folder by the recipient.
//                                   Unfortunately, only some of the mail services report such information, so usually
//                                   you won’t receive a lot of responses with such status.
//
//   ok_link_visited               – The message has been delivered, read, and one of the links has been clicked through.
//                                   It may change to ‘ok_unsubscribed’ or ‘ok_spam_folder’.
//
//   ok_unsubscribed               – The message has been delivered and read, but the user unsubscribed using the link
//                                   in the letter. The status is final.
//
//   err_blacklisted               – The message has been rejected due to blacklisting.
//
//   err_will_retry                – One or more delivery attempts have been unsuccessful, but attempts continue.
//                                   The status is not final.
//
//   err_resend                    – Actually, it is equivalent to err_will_retry with some minor internal features.
//
//   err_internal                  – Internal failure. The letter needs to be re-sent. The status is final.
//
//   err_user_unknown              – The address does not exist, delivery failed. The status is final.
//
//   err_user_inactive             – The address used to exist before, but now it has been disabled. Delivery failed.
//                                   The status is final.
//
//   err_mailbox_discarded         – The recipient’s mailbox has been deleted. The status is final.
//
//   err_mailbox_full              – The recipient’s mailbox is full. The status is final.
//
//   err_no_dns                    – There is no record or an incorrect record in DNS.
//
//   err_no_smtp                   – There is an entry in DNS, but there is no smtp server.
//
//   err_domain_inactive           – The domain does not accept mail or does not exist. The status is final.
//
//   err_destination_misconfigured – The domain does not accept mail due to incorrect settings on the recipient’s side,
//                                   and the server’s response contains information about a cause that can be fixed
//                                   (for example, an inoperative blacklist is used, etc.).
//
//   err_spam_rejected             – The message was rejected by the server as a spam.
//
//   err_too_large                 – The letter exceeds the size allowed by the recipient’s server. Also, the reason
//                                   may be a rejection of the letter by the recipient’s server due to an unacceptable
//                                   attachment type. For example, .exe.
//
//   err_giveup                    – This status is assigned to messages with the err_will_retry, err_resend statuses
//                                   after expiration of the retry period.
//
//   err_spam_removed              – Sending has been canceled because the campaign has been clocked as a spam.
//                                   The status is not final, it can be changed to not_sent, delayed
//                                   or err_spam_may_retry after negotiations with the recipient’s mail service.
//
//   err_spam_may_retry            – It is equivalent to err_spam_rejected, but you can re-send the message
//                                   by generating a new similar letter.
//
//   ok_spam_folder                – The letter has been delivered, but the recipient’s server placed it
//                                   in the Spam folder. The status is final.
//
//   err_delivery_failed           – Delivery failed due to other reasons. The status is final.
//
//   err_will_retry                – One or more delivery attempts have been unsuccessful, but attempts continue.
//                                   The status is not final.
//
//   err_skip_letter               – Sending has been canceled because the email address is not available
//                                   (except for cases of err_unsubscribed and err_not_allowed).
//
//   err_spam_skipped              – Sending has been canceled because the campaign has been blocked as a spam.
//                                   The result is not final, it can be changed to not_sent, delayed
//                                   or err_spam_may_retry after negotiations with the recipient’s mail service.
//
//   err_unsubscribed              – The letter has not been sent as the address to which it was tried to be sent
//                                   has previously unsubscribed. You can mark this address as unsubscribed
//                                   in your database and not send messages to it any more. The status is final.
//
//   err_src_invalid               – Invalid sender’s email address. It is used if the “invalid sender’s email”
//                                   was discovered not at the stage of accepting the task and checking the parameters,
//                                   but at a later stage when the things to be sent are checked in detail.
//                                   The status is final.
//
//   err_dest_invalid              – Invalid recipient’s email address. It is used if the “invalid recipient’s email”
//                                   was discovered not at the stage of accepting the task and checking the parameters,
//                                   but at a later stage when the things to be sent are checked in detail.
//                                   The status is final.
//
//   err_not_allowed               – Sending has been canceled because technical support staff blocked the campaign,
//                                   or because of the recipient address or your account is blocked.
//                                   The status is final.
//
//   err_over_quota                – Sending has been canceled due to insufficiency of funds on the account or due
//                                   to an excess of the tariff.
//
//   err_not_available             – The address to which you have tried to send the letter is not available
//                                   (i.e., previous sending to this address resulted in a response
//                                   like “the address does not exist” or “block for spam” from the server).
//                                   Theoretically, the availability of the address can be restored in a few days
//                                   or weeks, so you may not eliminate it completely from the list of potential
//                                   recipients. The status is final.
//
//   err_unreachable               – Sending has been canceled because the address is not available, but, in contrast
//                                   to the err_not_available status, the availability of the address
//                                   will not be restored. The status is final.
//
//   err_lost                      – The letter wasn’t sent due to inconsistency of its parts (for example,
//                                   a link to the image in attachments is transferred in the letter body,
//                                   but the image itself is not transferred in the attachments), or it has been lost
//                                   because of the failure on our side. The sender needs to re-send the letter
//                                   on his own, since the original has not been saved. The status is final.
//
//   skip_dup_unreachable          – The address is unavailable, sending failed. The status is final.
//
//   skip_dup_temp_unreachable     – The address is temporarily unavailable. Sending failed. The status is final.
//
//   skip_dup_mailbox_full         – The recipient’s mailbox is full. The status is final.
//
// See: https://www.unisender.com/en/support/api/messages/checkemail/
func (u *UniSender) CheckEmail(emailIDs ...int64) *messages.CheckEmailRequest {
	return messages.CheckEmail(u.request(), emailIDs...)
}

// CheckSMS returns request to check the delivery status of sms sent using the sendSMS method.
//
// Status possible options:
//   not_sent            – The message has not been sent yet, and is waiting to be sent. The status will
//                         be changed after sending.
//
//   ok_sent             – The message has been sent, but the delivery status is still unknown. The status is temporary
//                         and may change.
//
//   ok_delivered        – The message has been delivered. The status is final.
//
//   err_src_invalid     – The delivery is not possible, the sender is set incorrectly. The status is final.
//
//   err_dest_invalid    – The delivery is not possible, a wrong number is indicated. The status is final.
//
//   err_skip_letter     – The delivery is impossible because the status of the phone number was changed
//                         in the process of sending, or the phone number has been removed from the list,
//                         or the letter has been deleted. The status is final.
//
//   err_not_allowed     – The delivery is not possible, this communications service provider is not serviced.
//                         The status is final.
//
//   err_delivery_failed – The delivery failed — usually because of indication of a formally correct,
//                         but non-existent number, or because the phone is turned off. The status is final.
//
//   err_lost            – The message has been lost, and the sender needs to re-send the message on his own,
//                         since the original has not been saved. The status is final.
//
//   err_internal        – Internal failure. The message needs to be re-sent. The status is final.
//
// See: https://www.unisender.com/en/support/api/partners/check-sms/
func (u *UniSender) CheckSMS(smsID int64) *messages.CheckSMSRequest {
	return messages.CheckSMS(u.request(), smsID)
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
