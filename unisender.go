package unisender

import (
	"net/http"
	"sync"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/common"
	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/messages"
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

// GetCurrencyRates returns currency rates.
func (u *UniSender) GetCurrencyRates() *common.GetCurrencyRatesRequest {
	return common.GetCurrencyRates(u.request())
}

// CreateList creates a new contact list.
func (u *UniSender) CreateList(title string) *contacts.CreateListRequest {
	return contacts.CreateList(u.request(), title)
}

// GetLists returns all available campaign lists.
func (u *UniSender) GetLists() *contacts.GetListsRequest {
	return contacts.GetLists(u.request())
}

// UpdateList changes campaign list properties.
func (u *UniSender) UpdateList(listID int64, title string) *contacts.UpdateListRequest {
	return contacts.UpdateList(u.request(), listID, title)
}

// DeleteList removes a list.
func (u *UniSender) DeleteList(listID int64) *contacts.DeleteListRequest {
	return contacts.DeleteList(u.request(), listID)
}

// GetContact returns information about single contact.
func (u *UniSender) GetContact(email string) *contacts.GetContactRequest {
	return contacts.GetContact(u.request(), email)
}

// Subscribe subscribes the contact email or phone number to one or several lists.
func (u *UniSender) Subscribe() *contacts.SubscribeRequest {
	return contacts.Subscribe(u.request())
}

// Unsubscribe unsubscribes the contact email or phone number from one or several lists.
func (u *UniSender) Unsubscribe(contact string) *contacts.UnsubscribeRequest {
	return contacts.Unsubscribe(u.request(), contact)
}

// Exclude excludes the contact’s email or phone number from one or several lists.
func (u *UniSender) Exclude(contact string) *contacts.ExcludeRequest {
	return contacts.Exclude(u.request(), contact)
}

// ImportContacts imports contacts.
func (u *UniSender) ImportContacts(collection *contacts.ImportContactsCollection) *contacts.ImportContactsRequest {
	return contacts.ImportContacts(u.request(), collection)
}

// ExportContacts export of contact data from UniSender.
func (u *UniSender) ExportContacts() *contacts.ExportContactsRequest {
	return contacts.ExportContacts(u.request())
}

// IsContactInList checks whether the contact is in the specified user lists.
func (u *UniSender) IsContactInList(email string, listIDs ...int64) *contacts.IsContactInListRequest {
	return contacts.IsContactInList(u.request(), email, listIDs...)
}

// GetContactCount returns the contacts list size.
func (u *UniSender) GetContactCount(listID int64) *contacts.GetContactCountRequest {
	return contacts.GetContactCount(u.request(), listID)
}

// GetTotalContactsCount returns the contacts database size by the user login.
func (u *UniSender) GetTotalContactsCount(login string) *contacts.GetTotalContactsCountRequest {
	return contacts.GetTotalContactsCount(u.request(), login)
}

// CheckEmail returns request to check the delivery status of emails sent using.
func (u *UniSender) CheckEmail(emailIDs ...int64) *messages.CheckEmailRequest {
	return messages.CheckEmail(u.request(), emailIDs...)
}

// GetCheckedEmail returns request to check the delivery status of emails sent using.
func (u *UniSender) GetCheckedEmail(login string) *messages.GetCheckedEmailRequest {
	return messages.GetCheckedEmail(u.request(), login)
}

// ValidateSender returns request that sends a message to the email address with a link to confirm the address
// as the return address.
func (u *UniSender) ValidateSender(email string) *messages.ValidateSenderRequest {
	return messages.ValidateSender(u.request(), email)
}

// GetSenderDomainList returns information about domains.
func (u *UniSender) GetSenderDomainList(login string) *messages.GetSenderDomainListRequest {
	return messages.GetSenderDomainList(u.request(), login)
}

// SetSenderDomain register the domain in the list.
func (u *UniSender) SetSenderDomain(login, domain string) *messages.SetSenderDomainRequest {
	return messages.SetSenderDomain(u.request(), login, domain)
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
