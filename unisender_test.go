package unisender_test

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/alexeyco/unisender"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//// RoundTripFunc .
//type RoundTripFunc func(req *http.Request) *http.Response
//
//// RoundTrip .
//func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
//	return f(req), nil
//}

type roundTrip struct {
	Before     func(url *url.URL)
	StatusCode func() int
	Body       func() string
}

func (r roundTrip) RoundTrip(req *http.Request) (*http.Response, error) {
	if r.Before != nil {
		r.Before(req.URL)
	}

	statusCode := http.StatusOK
	if r.StatusCode != nil {
		statusCode = r.StatusCode()
	}

	var body string
	if r.Body != nil {
		body = r.Body()
	}

	return &http.Response{
		StatusCode: statusCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
		Header:     make(http.Header),
	}, nil
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(
		time.Now().UnixNano(),
	),
)

const (
	apiKeyLength = 32
	charset      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func randomAPIKey() string {
	return randomString(apiKeyLength)
}

func randomString(n int) string {
	l := len(charset)
	b := make([]byte, n)

	for i := range b {
		b[i] = charset[seededRand.Intn(l)]
	}

	return string(b)
}

func randomInt64(from, to int) int64 {
	return int64(from + seededRand.Intn(to-from))
}

func randomLanguage() (language string) {
	languages := []string{
		unisender.LanguageEnglish,
		unisender.LanguageRussian,
		unisender.LanguageItalian,
	}

	return languages[seededRand.Intn(len(languages)-1)]
}

func TestUnisender(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unisender test suite")
}
