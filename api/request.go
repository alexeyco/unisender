package api

import (
	"net/http"
	"net/url"
)

const endpoint = "https://api.unisender.com/"

type Request struct {
	client    *http.Client
	logger    Logger
	url       string
	arguments url.Values
}

func (r *Request) Add(key, value string) *Request {
	r.arguments.Add(key, value)
	return r
}

func (r *Request) Execute(v interface{}) (err error) {
	return
}

func NewRequest(client *http.Client, language string) *Request {
	return &Request{
		client:    client,
		url:       endpoint + language + "/api/",
		arguments: url.Values{},
	}
}
