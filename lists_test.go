package unisender_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender"
)

func TestUnisender_GetLists(t *testing.T) {
	expectedApiKey := randomString(32)
	var requestedApiKey string

	language := randomLanguage()

	expectedUrl := fmt.Sprintf("https://api.unisender.com/%s/api/getLists", language)
	var requestedUrl string

	id := randomInt64(999, 99999)
	title := randomString(64)

	client := &http.Client{
		Transport: roundTrip{
			Before: func(req *http.Request) {
				requestedUrl = req.URL.String()
				requestedApiKey = req.PostFormValue("api_key")
			},
			Body: func() string {
				return fmt.Sprintf(`{"result":[{"id":%d,"title":"%s"}]}`, id, title)
			},
		},
	}

	usndr := unisender.New(expectedApiKey)
	usndr.SetLanguage(language)
	usndr.SetClient(client)

	lists, err := usndr.GetLists()

	if expectedUrl != requestedUrl {
		t.Fatalf(`Request URL should be "%s", "%s" given`, expectedUrl, requestedUrl)
	}

	if expectedApiKey != requestedApiKey {
		t.Fatalf(`API key should be "%s", "%s" given`, expectedApiKey, requestedApiKey)
	}

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	l := len(lists)
	if l != 1 {
		t.Fatalf("Lists length should be %d, %d given", 1, l)
	}

	if lists[0].ID != id {
		t.Fatalf("List ID should be %d, %d given", id, lists[0].ID)
	}

	if lists[0].Title != title {
		t.Fatalf(`List title should be "%s", "%s" given`, title, lists[0].Title)
	}
}
