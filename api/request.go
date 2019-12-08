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

// SetLogger sets request logger.
func (r *Request) SetLogger(logger Logger) *Request {
	r.logger = logger
	return r
}

// Add adds request argument.
func (r *Request) Add(key, value string) *Request {
	r.arguments.Add(key, value)
	return r
}

// Execute executes request and map response to specified value.
func (r *Request) Execute(method string, v interface{}) (err error) {
	u := r.url + method
	r.logRequest("POST", u, r.arguments)

	var resp *http.Response
	if resp, err = r.client.PostForm(u, r.arguments); err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		//r.logResponse()
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

func (r *Request) logRequest(method, url string, values url.Values) {
	if r.logger == nil {
		return
	}

	r.logger.LogRequest(method, url, values)
}

func (r *Request) logResponse() {
	if r.logger == nil {
		return
	}

}

// NewRequest returns new API request.
func NewRequest(client *http.Client, language string) *Request {
	return &Request{
		client:    client,
		url:       Endpoint + language + "/api/",
		arguments: url.Values{},
	}
}
