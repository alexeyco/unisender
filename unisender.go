package unisender

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

type requestResult struct {
	Result interface{} `json:"result"`
}

func (u *Unisender) request(method string, data url.Values, v interface{}) (err error) {
	uri := fmt.Sprintf(apiEndpointPattern, u.language, method)
	data.Add("api_key", u.apiKey)

	var resp *http.Response
	if resp, err = u.client.PostForm(uri, data); err != nil {
		return
	}

	result := requestResult{
		Result: v,
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	v = result.Result

	return
}

// New returns new Unisender client
func New(apiKey string) *Unisender {
	return &Unisender{
		language: DefaultLanguage,
		apiKey:   apiKey,
		client:   http.DefaultClient,
	}
}
