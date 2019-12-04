package unisender

import (
	"net/http"
	"sync"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/lists"
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
func (u *UniSender) CreateList(title string) lists.CreateListRequest {
	return lists.CreateList(u.request(), title)
}

// GetLists returns all available campaign lists.
func (u *UniSender) GetLists() lists.GetListsRequest {
	return lists.GetLists(u.request())
}

// UpdateList changes campaign list properties.
func (u *UniSender) UpdateList(listID int64, title string) lists.UpdateListRequest {
	return lists.UpdateList(u.request(), listID, title)
}

// UpdateList removes a list.
func (u *UniSender) DeleteList(listID int64) lists.DeleteListRequest {
	return lists.DeleteList(u.request(), listID)
}

// GetContact returns information about single contact.
// See https://www.unisender.com/en/support/api/contacts/getcontact/
func (u *UniSender) GetContact() {

}

// Subscribe subscribes the contact email or phone number to one or several lists.
// See https://www.unisender.com/en/support/api/contacts/subscribe/
func (u *UniSender) Subscribe() {

}

// Unsubscribe unsubscribes the contact email or phone number from one or several lists.
// See https://www.unisender.com/en/support/api/partners/unsubscribe/
func (u *UniSender) Unsubscribe() {

}

// Exclude excludes the contact’s email or phone number from one or several lists.
// See https://www.unisender.com/en/support/api/partners/exclude/
func (u *UniSender) Exclude() {

}

// ImportContacts
// See https://www.unisender.com/en/support/api/contacts/importcontacts/
func (u *UniSender) ImportContacts() {

}

// ExportContacts
// See https://www.unisender.com/en/support/api/contacts/exportcontacts/
func (u *UniSender) ExportContacts() {

}

// IsContactInList checks whether the contact is in the specified user lists
// See https://www.unisender.com/en/support/api/contacts/iscontactinlist/
func (u *UniSender) IsContactInList() {

}

// GetContactCount returns the contacts list size.
// See https://www.unisender.com/en/support/api/contacts/getcontactcount/
func (u *UniSender) GetContactCount() {

}

// GetTotalContactCount returns the contacts database size by the user login.
// See https://www.unisender.com/en/support/api/partners/gettotalcontactscount/
func (u *UniSender) GetTotalContactCount() {

}

func (u *UniSender) request() *api.Request {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return api.NewRequest(u.client, u.language).
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
