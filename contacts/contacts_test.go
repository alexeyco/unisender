package contacts_test

import (
	"math/rand"
	"net/http"
	"time"

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

func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(min, max int) string {
	length := randomInt(min, max)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func randomTime(min, max int) time.Time {
	return time.Now().AddDate(0, 0, -randomInt(min, max))
}
