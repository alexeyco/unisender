package unisender_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender"
)

func TestUnisender_GetLists(t *testing.T) {
	apiKey := randomAPIKey()
	language := randomLanguage()

	expectedUrl := "https://api.unisender.com/" + language + "/api/getLists"

	id := randomInt64(999, 99999)
	title := randomString(64)

	client := &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			if req.URL.String() != expectedUrl {
				t.Fatalf(`Request URL should be "%s", but "%s" given`, expectedUrl, req.URL.String())
			}

			if req.Method != http.MethodPost {
				t.Fatalf(`Request method should be "%s", but "%s" given`, http.MethodPost, req.Method)
			}

			body := fmt.Sprintf(`{"result":[{"id":%d,"title":%s}]}`, id, title)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
				Header:     make(http.Header),
			}
		}),
	}

	usndr := unisender.New(apiKey)
	usndr.SetLanguage(language)
	usndr.SetClient(client)

	lists, err := usndr.GetLists()

	if err != nil {
		t.Fatalf(`Error should be nil, but "%s" given`, err.Error())
	}

	if len(lists) != 1 {
		t.Fatalf(`List length should be 1, but %d given`, len(lists))
	}

	if lists[0].ID != id {
		t.Fatalf(`List ID should be %d, but %d given`, id, lists[0].ID)
	}

	if lists[0].Title != title {
		t.Fatalf(`List title should be %s, but %s given`, title, lists[0].Title)
	}
}
