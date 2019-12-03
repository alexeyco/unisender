package unisender_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/alexeyco/unisender"
)

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
