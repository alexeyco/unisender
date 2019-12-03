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

// SetClient sets custom http.Client to UniSender client
func (u *Unisender) SetClient(client *http.Client) {
	u.client = client
}

// request makes request to UniSender API and parses response
func (u *Unisender) request(method string, v interface{}, options ...Option) (err error) {
	uri := fmt.Sprintf(apiEndpointPattern, u.language, method)

	data := url.Values{}
	data.Add("api_key", u.apiKey)

	for _, o := range options {
		data.Add(o.name, o.value)
	}

	var resp *http.Response
	if resp, err = u.client.PostForm(uri, data); err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return ErrWrongStatusCode
	}

	response := Response{
		Result: v,
	}

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return
	}

	if err = response.Err(); err != nil {
		return
	}

	v = response.Result

	return
}

// New returns new UniSender client
func New(apiKey string) *Unisender {
	return &Unisender{
		language: DefaultLanguage,
		apiKey:   apiKey,
		client:   http.DefaultClient,
	}
}
