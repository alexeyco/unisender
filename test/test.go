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

// NewRequest returns test API request.
func NewRequest(fn func(req *http.Request) (res *http.Response, err error)) *api.Request {
	return api.NewRequest(NewClient(fn), "en")
}

// NewClient returns test http client.
func NewClient(fn func(req *http.Request) (res *http.Response, err error)) *http.Client {
	return &http.Client{
		Transport: &roundTripper{
			fn: fn,
		},
	}
}

// RandomInt returns an int in the given interval.
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomInt64 returns an int64 in the given interval.
func RandomInt64(min, max int) int64 {
	return int64(RandomInt(min, max))
}

// RandomTime random time.Time.
func RandomTime(min, max int) time.Time {
	return time.Now().
		AddDate(0, 0, -RandomInt(min, max)).
		Round(time.Second).
		UTC()
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString returns random string with random length in given interval.
func RandomString(min, max int) string {
	length := RandomInt(min, max)

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// RandomInt64Slice returns []int64 slice with random length in given interval.
func RandomInt64Slice(min, max int) (slice []int64) {
	l := RandomInt(min, max)
	for i := 0; i < l; i++ {
		slice = append(slice, RandomInt64(9999, 999999))
	}

	return slice
}

// RandomStringSlice returns []string slice with random length in given interval.
func RandomStringSlice(min, max int) []string {
	l := RandomInt(min, max)
	slice := make([]string, l)
	for i := 0; i < l; i++ {
		slice[i] = RandomString(12, 32)
	}

	return slice
}
