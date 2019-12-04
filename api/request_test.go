package api_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/api"
)

type roundTripper struct {
	fn func(req *http.Request) (res *http.Response, err error)
}

func (r *roundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	return r.fn(req)
}

func newClient(fn func(req *http.Request) (res *http.Response, err error)) *http.Client {
	return &http.Client{
		Transport: &roundTripper{
			fn: fn,
		},
	}
}

func TestRequest_Add(t *testing.T) {
	method := "method"
	language := "some-language"

	var argRequested string

	c := newClient(func(req *http.Request) (res *http.Response, err error) {
		argRequested = req.PostFormValue("arg")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	req := api.NewRequest(c, language)

	argExpected := "some-value"
	err := req.Add("arg", argExpected).
		Execute(method, nil)

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if argRequested != argExpected {
		t.Fatalf(`Arg value should be "%s", "%s" given`, argExpected, argRequested)
	}
}

func TestRequest_Execute_Error(t *testing.T) {
	method := "method"
	language := "some-language"

	c := newClient(func(req *http.Request) (res *http.Response, err error) {
		return &http.Response{
			StatusCode: http.StatusForbidden,
		}, nil
	})

	req := api.NewRequest(c, language)
	err := req.Execute(method, nil)

	if err == nil {
		t.Fatalf(`Error should be "%s", nil given`, api.ErrWrongStatusCode.Error())
	}

	if err != api.ErrWrongStatusCode {
		t.Fatalf(`Error should be "%s", "%s" given`, api.ErrWrongStatusCode.Error(), err.Error())
	}
}

type testResponse struct {
	Foo string `json:"foo"`
}

func TestRequest_Execute_Ok(t *testing.T) {
	method := "method"
	language := "some-language"

	var urlRequested string
	var methodRequested string

	c := newClient(func(req *http.Request) (res *http.Response, err error) {
		methodRequested = req.Method
		urlRequested = req.URL.String()

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"result":{"foo":"bar"}}`)),
		}, nil
	})

	req := api.NewRequest(c, language)

	var res testResponse
	err := req.Execute(method, &res)

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if methodRequested != http.MethodPost {
		t.Fatalf(`URL should be "%s", "%s" given`, http.MethodPost, methodRequested)
	}

	urlExpected := api.Endpoint + language + "/api/" + method
	if urlRequested != urlExpected {
		t.Fatalf(`URL should be "%s", "%s" given`, urlExpected, urlRequested)
	}

	if res.Foo != "bar" {
		t.Fatalf(`Parsed response param value should be "%s", "%s given"`, "bar", res.Foo)
	}
}
