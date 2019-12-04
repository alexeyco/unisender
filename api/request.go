package api

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Endpoint API endpoint.
const Endpoint = "https://api.unisender.com/"

// Request API request.
type Request struct {
	client    *http.Client
	logger    Logger
	url       string
	arguments url.Values
}

// Add adds request argument.
func (r *Request) Add(key, value string) *Request {
	r.arguments.Add(key, value)
	return r
}

// Execute executes request and map response to specified value.
func (r *Request) Execute(method string, v interface{}) (err error) {
	var resp *http.Response
	if resp, err = r.client.PostForm(r.url+method, r.arguments); err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return ErrWrongStatusCode
	}

	if v == nil {
		return
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

// NewRequest returns new API request.
func NewRequest(client *http.Client, language string) *Request {
	return &Request{
		client:    client,
		url:       Endpoint + language + "/api/",
		arguments: url.Values{},
	}
}
