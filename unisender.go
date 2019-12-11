package unisender

import (
	"net/http"
	"sync"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/campaigns"
	"github.com/alexeyco/unisender/common"
	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/lists"
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

// Exclude returns request, that excludes the contact’s email or phone number from one or several lists.
// In contrast to the unsubscribe method, it does not mark the contact as "unsubscribed", and it can be included
// in the list again later using the subscribe method.
//
// See https://www.unisender.com/en/support/api/partners/exclude/
func (u *UniSender) Exclude(contact string) *contacts.ExcludeRequest {
	return contacts.Exclude(u.request(), contact)
}

// ExportContacts returns request, that exports of contact data from UniSender.
//
// See https://www.unisender.com/en/support/api/contacts/exportcontacts/
func (u *UniSender) ExportContacts() *contacts.ExportContactsRequest {
	return contacts.ExportContacts(u.request())
}

// GetContact returns request to getting information about a contact (one contact only).
//
// See https://www.unisender.com/en/support/api/contacts/getcontact/
func (u *UniSender) GetContact(email string) *contacts.GetContactRequest {
	return contacts.GetContact(u.request(), email)
}

// GetContactCount returns request for count contacts by specified conditions.
//
// See https://www.unisender.com/en/support/api/contacts/getcontactcount/
func (u *UniSender) GetContactCount(listID int64) *contacts.GetContactCountRequest {
	return contacts.GetContactCount(u.request(), listID)
}

// GetContactFieldValues returns request to obtain the values of the specified additional contact fields.
//
// See: https://www.unisender.com/en/support/api/partners/getcontactfieldvalues/
func (u *UniSender) GetContactFieldValues(email string, fieldIDs ...int64) *contacts.GetContactFieldValuesRequest {
	return contacts.GetContactFieldValues(u.request(), email, fieldIDs...)
}

// GetFields returns request to get the list of user fields.
//
// See: https://www.unisender.com/en/support/api/partners/getfields/
func (u *UniSender) GetFields() *contacts.GetFieldsRequest {
	return contacts.GetFields(u.request())
}

// GetFields returns request to get the list of user tags.
//
// See: https://www.unisender.com/en/support/api/partners/gettags/
func (u *UniSender) GetTags() *contacts.GetTagsRequest {
	return contacts.GetTags(u.request())
}

// GetTotalContactsCount returns request that counts contacts database size by the user login.
//
// See: https://www.unisender.com/en/support/api/partners/gettotalcontactscount/
func (u *UniSender) GetTotalContactsCount(login string) *contacts.GetTotalContactsCountRequest {
	return contacts.GetTotalContactsCount(u.request(), login)
}

// ImportContacts returns request to bulk import of contacts. It can also be used for periodic synchronization
// with the contact database stored on your own server (see also the description of the exportContacts method).
// You can import data of no more than 500 contacts per call. Larger lists must be imported in a few calls.
//
// If there are new addresses among the signed e-mail addresses, then by default they receive the status "new".
//
// Technical restrictions: the maximum number of user fields is 50. The timeout per call is 30 seconds from the moment
// the request is completely transmitted to the server. If no response is received after the timeout,
// then it is recommended to make up to two retries, and if there is no answer again, then contact technical support.
//
// See https://www.unisender.com/en/support/api/contacts/importcontacts/
func (u *UniSender) ImportContacts(collection *contacts.ImportContactsCollection) *contacts.ImportContactsRequest {
	return contacts.ImportContacts(u.request(), collection)
}

// IsContactInList returns request to check whether the contact is in the specified user lists.
//
// See https://www.unisender.com/en/support/api/contacts/iscontactinlist/
func (u *UniSender) IsContactInList(email string, listIDs ...int64) *contacts.IsContactInListRequest {
	return contacts.IsContactInList(u.request(), email, listIDs...)
}

// Subscribe returns request that adds contacts (email address and/or mobile phone) of a contact to one
// or several lists, and also allows you to add/change the values ​​of additional fields and labels.
//
// See https://www.unisender.com/en/support/api/contacts/subscribe/
func (u *UniSender) Subscribe() *contacts.SubscribeRequest {
	return contacts.Subscribe(u.request())
}

// Unsubscribe returns request that unsubscribes the contact email or phone number from one or several lists.
// In contrast to the exclude method, it does not exclude a contact from the lists, but marks the contact
// as "unsubscribed". It is impossible to restore the «active» status through API – it is only the contact
// who can do this by clicking on the activation link in the letter.
//
// See https://www.unisender.com/en/support/api/partners/unsubscribe/
func (u *UniSender) Unsubscribe(contact string) *contacts.UnsubscribeRequest {
	return contacts.Unsubscribe(u.request(), contact)
}

// UpdateField returns request to change user field parameters.
//
// See: https://www.unisender.com/en/support/api/partners/updatefield/
func (u *UniSender) UpdateField(fieldID int64, name string) *contacts.UpdateFieldRequest {
	return contacts.UpdateField(u.request(), fieldID, name)
}

// CreateList returns request to create a new contact list.
//
// See https://www.unisender.com/en/support/api/partners/createlist/
func (u *UniSender) CreateList(title string) *lists.CreateListRequest {
	return lists.CreateList(u.request(), title)
}

// DeleteList returns request to delete a list.
//
// See https://www.unisender.com/en/support/api/partners/deletelist/
func (u *UniSender) DeleteList(listID int64) *lists.DeleteListRequest {
	return lists.DeleteList(u.request(), listID)
}

// GetLists returns request to get the list of all available campaign lists.
//
// See https://www.unisender.com/en/support/api/partners/getlists/
func (u *UniSender) GetLists() *lists.GetListsRequest {
	return lists.GetLists(u.request())
}

// UpdateList returns request to change campaign list properties.
//
// See https://www.unisender.com/en/support/api/partners/updatelist/
func (u *UniSender) UpdateList(listID int64, title string) *lists.UpdateListRequest {
	return lists.UpdateList(u.request(), listID, title)
}

// UpdateOptInEmail updates OptInEmail text. Each campaign list has the attached text of the invitation
// to subscribe and confirm the email that is sent to the contact to confirm the campaign. If the contact subscribes
// to several lists at once, a letter with the text of the first list will be sent. The text of the letter
// can be changed using the updateOptInEmail method. The text must include at least one link with
// the attribute href=”{{ConfirmUrl}}”.
//
// See https://www.unisender.com/en/support/api/partners/updateoptinemail/
func (u *UniSender) UpdateOptInEmail(listID int64) *lists.UpdateOptInEmailRequest {
	return lists.UpdateOptInEmail(u.request(), listID)
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

// CreateEmailMessage returns request to create an email without sending it.
// The email is sent using another method — createCampaign.
//
// See: https://www.unisender.com/en/support/api/partners/createemailmessage/
func (u *UniSender) CreateEmailMessage(listID int64) *messages.CreateEmailMessageRequest {
	return messages.CreateEmailMessage(u.request(), listID)
}

// CreateEmailTemplate returns request to create an email template for a mass campaign.
//
// See: https://www.unisender.com/en/support/api/partners/createemailmessage/
func (u *UniSender) CreateEmailTemplate(title string) *messages.CreateEmailTemplateRequest {
	return messages.CreateEmailTemplate(u.request(), title)
}

// CreateSMSMessage returns request to create an sms without sending it.
// The sms is sent using another method — createCampaign.
//
// See: https://www.unisender.com/en/support/api/messages/createsmsmessage/
func (u *UniSender) CreateSMSMessage(sender string) *messages.CreateSMSMessageRequest {
	return messages.CreateSMSMessage(u.request(), sender)
}

// DeleteMessage returns request to delete a message.
//
// See: https://www.unisender.com/en/support/api/partners/deletemessage/
func (u *UniSender) DeleteMessage(messageID int64) *messages.DeleteMessageRequest {
	return messages.DeleteMessage(u.request(), messageID)
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
