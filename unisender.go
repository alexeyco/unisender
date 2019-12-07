package unisender

import (
	"net/http"
	"sync"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts"
)

type UniSender struct {
	apiKey   string
	language string
	client   *http.Client
	mu       sync.RWMutex
}

// SetLanguage sets API response language
func (u *UniSender) SetLanguage(language string) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.language = language
}

// SetClient sets custom http.Request to UniSender client
func (u *UniSender) SetClient(client *http.Client) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.client = client
}

// CreateList creates a new contact list.
func (u *UniSender) CreateList(title string) contacts.CreateListRequest {
	return contacts.CreateList(u.request(), title)
}

// GetLists returns all available campaign lists.
func (u *UniSender) GetLists() contacts.GetListsRequest {
	return contacts.GetLists(u.request())
}

// UpdateList changes campaign list properties.
func (u *UniSender) UpdateList(listID int64, title string) contacts.UpdateListRequest {
	return contacts.UpdateList(u.request(), listID, title)
}

// UpdateList removes a list.
func (u *UniSender) DeleteList(listID int64) contacts.DeleteListRequest {
	return contacts.DeleteList(u.request(), listID)
}

// GetContact returns information about single contact.
func (u *UniSender) GetContact(email string) contacts.GetContactRequest {
	return contacts.GetContact(u.request(), email)
}

// Subscribe subscribes the contact email or phone number to one or several lists.
func (u *UniSender) Subscribe() contacts.SubscribeRequest {
	return contacts.Subscribe(u.request())
}

// Unsubscribe unsubscribes the contact email or phone number from one or several lists.
func (u *UniSender) Unsubscribe(contact string) contacts.UnsubscribeRequest {
	return contacts.Unsubscribe(u.request(), contact)
}

// Exclude excludes the contact’s email or phone number from one or several lists.
func (u *UniSender) Exclude(contact string) contacts.ExcludeRequest {
	return contacts.Exclude(u.request(), contact)
}

// ImportContacts imports contacts.
func (u *UniSender) ImportContacts(collection *contacts.Collection) contacts.ImportContactsRequest {
	return contacts.ImportContacts(u.request(), collection)
}

// ExportContacts export of contact data from UniSender.
func (u *UniSender) ExportContacts() contacts.ExportContactsRequest {
	return contacts.ExportContacts(u.request())
}

// IsContactInList checks whether the contact is in the specified user lists
func (u *UniSender) IsContactInList(email string, listIDs ...int64) contacts.IsContactInListRequest {
	return contacts.IsContactInList(u.request(), email, listIDs...)
}

// GetContactCount returns the contacts list size.
func (u *UniSender) GetContactCount(listID int64) contacts.GetContactCountRequest {
	return contacts.GetContactCount(u.request(), listID)
}

// GetTotalContactsCount returns the contacts database size by the user login.
func (u *UniSender) GetTotalContactsCount(login string) contacts.GetTotalContactsCountRequest {
	return contacts.GetTotalContactsCount(u.request(), login)
}

func (u *UniSender) request() *api.Request {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return api.NewRequest(u.client, u.language).
		Add("format", "json").
		Add("api_key", u.apiKey)
}

// New returns new UniSender API client
func New(apiKey string) *UniSender {
	return &UniSender{
		apiKey:   apiKey,
		language: api.DefaultLanguage,
		client:   http.DefaultClient,
		mu:       sync.RWMutex{},
	}
}
