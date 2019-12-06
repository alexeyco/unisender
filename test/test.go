package test

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/alexeyco/unisender/api"
)

type roundTripper struct {
	fn func(req *http.Request) (res *http.Response, err error)
}

func (r *roundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	return r.fn(req)
}

// NewRequest returns testing API request.
func NewRequest(fn func(req *http.Request) (res *http.Response, err error)) *api.Request {
	return api.NewRequest(NewClient(fn), api.DefaultLanguage)
}

// NewClient returns testing http client.
func NewClient(fn func(req *http.Request) (res *http.Response, err error)) *http.Client {
	return &http.Client{
		Transport: &roundTripper{
			fn: fn,
		},
	}
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomInt64(min, max int) int64 {
	return int64(RandomInt(min, max))
}

func RandomTime(min, max int) time.Time {
	return time.Now().AddDate(0, 0, -RandomInt(min, max))
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(min, max int) string {
	length := RandomInt(min, max)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func RandomInt64Slice(min, max int) (slice []int64) {
	l := RandomInt(min, max)
	for i := 0; i < l; i++ {
		slice = append(slice, RandomInt64(9999, 999999))
	}

	return slice
}

func RandomStringSlice(min, max int) []string {
	l := RandomInt(min, max)
	slice := make([]string, l)
	for i := 0; i < l; i++ {
		slice[i] = RandomString(12, 32)
	}

	return slice
}
