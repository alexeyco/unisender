package unisender_test

import (
	"github.com/alexeyco/unisender"
	"math/rand"
	"net/http"
	"time"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
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
