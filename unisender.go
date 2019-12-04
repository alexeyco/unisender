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

func (u *UniSender) request() *api.Request {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return api.NewRequest(u.client, u.language).
		Add("api_key", u.apiKey)
}

// New returns new UniSender client
func New(apiKey string) *UniSender {
	return &UniSender{
		apiKey:   apiKey,
		language: api.DefaultLanguage,
		client:   http.DefaultClient,
		mu:       sync.RWMutex{},
	}
}
