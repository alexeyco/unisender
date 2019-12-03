package unisender_test

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/alexeyco/unisender"
)

type roundTrip struct {
	Before     func(req *http.Request)
	StatusCode func() int
	Body       func() string
}

func (r roundTrip) RoundTrip(req *http.Request) (*http.Response, error) {
	if r.Before != nil {
		r.Before(req)
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

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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

func TestUnisender_WrongStatusCode(t *testing.T) {
	apiKey := randomString(32)
	language := randomLanguage()

	client := &http.Client{
		Transport: roundTrip{
			StatusCode: func() int {
				return http.StatusForbidden
			},
			Body: func() string {
				return `{"Response":[]}`
			},
		},
	}

	usndr := unisender.New(apiKey)
	usndr.SetLanguage(language)
	usndr.SetClient(client)

	_, err := usndr.GetLists()

	if err == nil {
		t.Fatalf(`Error should be %s, nil given`, unisender.ErrWrongStatusCode.Error())
	}
}
func TestUnisender_Error(t *testing.T) {
	apiKey := randomString(32)
	language := randomLanguage()

	client := &http.Client{
		Transport: roundTrip{
			Body: func() string {
				return `{"Err":"Access Denied.","code":"access_denied"}`
			},
		},
	}

	usndr := unisender.New(apiKey)
	usndr.SetLanguage(language)
	usndr.SetClient(client)

	_, err := usndr.GetLists()

	if err == nil {
		t.Fatalf(`Error should be "%s", nil given`, unisender.ErrAccessDenied.Error())
	}

	if err != unisender.ErrAccessDenied {
		t.Fatalf(`Error should be "%s", %s given`, unisender.ErrAccessDenied.Error(), err.Error())
	}
}
