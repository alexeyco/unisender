package unisender_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/alexeyco/unisender"
)

func TestUnisender_GetLists(t *testing.T) {
	apiKey := randomString(32)
	language := randomLanguage()

	expectedUrl := fmt.Sprintf("https://api.unisender.com/%s/api/getLists", language)
	var requestedUrl string

	id := randomInt64(999, 99999)
	title := randomString(64)

	client := &http.Client{
		Transport: roundTrip{
			Before: func(url *url.URL) {
				requestedUrl = url.String()
			},
			Body: func() string {
				return fmt.Sprintf(`{"result":[{"id":%d,"title":"%s"}]}`, id, title)
			},
		},
	}

	usndr := unisender.New(apiKey)
	usndr.SetLanguage(language)
	usndr.SetClient(client)

	lists, err := usndr.GetLists()

	if expectedUrl != requestedUrl {
		t.Errorf("Request URL should be %s, %s requested", expectedUrl, requestedUrl)
	}

	if err != nil {
		t.Errorf("Error should be nil, %s given", err.Error())
	}

	l := len(lists)
	if l != 1 {
		t.Errorf("Lists length should be %d, %d given", 1, l)
	}

	if lists[0].ID != id {
		t.Errorf("List ID should be %d, %d given", id, lists[0].ID)
	}

	if lists[0].Title != title {
		t.Errorf(`List title should be "%s", "%s" given`, title, lists[0].Title)
	}
}
