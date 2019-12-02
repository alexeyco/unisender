package unisender

import (
	"net/http"
)

const apiEndpointPattern = "https://api.unisender.com/%s/api/%s"

// Unisender api client
type Unisender struct {
	language string
	apiKey   string
	client   *http.Client
}

// SetLanguage sets API response language
func (u *Unisender) SetLanguage(language string) {
	u.language = language
}

// SetClient sets custom http.Client to Unisender client
func (u *Unisender) SetClient(client *http.Client) {
	u.client = client
}

// New returns new Unisender client
func New(apiKey string) *Unisender {
	return &Unisender{
		language: DefaultLanguage,
		apiKey:   apiKey,
		client:   http.DefaultClient,
	}
}
