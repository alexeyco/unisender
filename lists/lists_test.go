package lists_test

import (
	"net/http"

	"github.com/alexeyco/unisender/api"
)

func newRequest(fn func(req *http.Request) (res *http.Response, err error)) *api.Request {
	return api.NewRequest(&http.Client{
		Transport: &roundTripper{
			fn: fn,
		},
	}, api.DefaultLanguage)
}

type roundTripper struct {
	fn func(req *http.Request) (res *http.Response, err error)
}

func (r *roundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	return r.fn(req)
}
