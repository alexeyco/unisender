package unisender_test

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender"
)

func TestUniSender_GetLists(t *testing.T) {
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

func TestUniSender_CreateList(t *testing.T) {
	expectedApiKey := randomString(32)
	var requestedApiKey string

	language := randomLanguage()

	expectedUrl := fmt.Sprintf("https://api.unisender.com/%s/api/createList", language)
	var requestedUrl string

	var requestedTitle string
	var requestedBeforeSubscribeUrl string
	var requestedAfterSubscribeUrl string

	expectedID := randomInt64(999, 99999)
	expectedTitle := randomString(64)

	client := &http.Client{
		Transport: roundTrip{
			Before: func(req *http.Request) {
				requestedUrl = req.URL.String()

				requestedApiKey = req.PostFormValue("api_key")

				requestedTitle = req.PostFormValue("title")
				requestedBeforeSubscribeUrl = req.PostFormValue("before_subscribe_url")
				requestedAfterSubscribeUrl = req.PostFormValue("after_subscribe_url")
			},
			Body: func() string {
				return fmt.Sprintf(`{"result":{"id":%d}}`, expectedID)
			},
		},
	}

	usndr := unisender.New(expectedApiKey)
	usndr.SetLanguage(language)
	usndr.SetClient(client)

	expectedBeforeSubscribeUrl := "https://before-subscribe.url"
	expectedAfterSubscribeUrl := "https://after-subscribe.url"

	createdID, err := usndr.CreateList(
		expectedTitle,
		unisender.OptionBeforeSubscribeUrl(expectedBeforeSubscribeUrl),
		unisender.OptionAfterSubscribeUrl(expectedAfterSubscribeUrl),
	)

	if expectedUrl != requestedUrl {
		t.Fatalf(`Request URL should be "%s", "%s" given`, expectedUrl, requestedUrl)
	}

	if expectedApiKey != requestedApiKey {
		t.Fatalf(`API key should be "%s", "%s" given`, expectedApiKey, requestedApiKey)
	}

	if expectedTitle != requestedTitle {
		t.Fatalf(`Title should be "%s", "%s" given`, expectedTitle, requestedTitle)
	}

	if expectedBeforeSubscribeUrl != requestedBeforeSubscribeUrl {
		t.Fatalf(`Param "before_subscribe_url" should be "%s", "%s" given`, expectedBeforeSubscribeUrl, requestedBeforeSubscribeUrl)
	}

	if expectedAfterSubscribeUrl != requestedAfterSubscribeUrl {
		t.Fatalf(`Param "after_subscribe_url" should be "%s", "%s" given`, expectedAfterSubscribeUrl, requestedAfterSubscribeUrl)
	}

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedID != createdID {
		t.Fatalf(`ID should be %d, %d given`, expectedID, createdID)
	}
}

func TestUniSender_UpdateList(t *testing.T) {
	expectedApiKey := randomString(32)
	var requestedApiKey string

	language := randomLanguage()

	expectedUrl := fmt.Sprintf("https://api.unisender.com/%s/api/updateList", language)
	var requestedUrl string

	var requestedID int64
	var requestedTitle string
	var requestedBeforeSubscribeUrl string
	var requestedAfterSubscribeUrl string

	expectedID := randomInt64(999, 99999)
	expectedTitle := randomString(64)

	client := &http.Client{
		Transport: roundTrip{
			Before: func(req *http.Request) {
				requestedUrl = req.URL.String()

				requestedApiKey = req.PostFormValue("api_key")

				requestedID, _ = strconv.ParseInt(req.PostFormValue("list_id"), 10, 64)
				requestedTitle = req.PostFormValue("title")
				requestedBeforeSubscribeUrl = req.PostFormValue("before_subscribe_url")
				requestedAfterSubscribeUrl = req.PostFormValue("after_subscribe_url")
			},
		},
	}

	usndr := unisender.New(expectedApiKey)
	usndr.SetLanguage(language)
	usndr.SetClient(client)

	expectedBeforeSubscribeUrl := "https://before-subscribe.url"
	expectedAfterSubscribeUrl := "https://after-subscribe.url"

	err := usndr.UpdateList(
		expectedID,
		expectedTitle,
		unisender.OptionBeforeSubscribeUrl(expectedBeforeSubscribeUrl),
		unisender.OptionAfterSubscribeUrl(expectedAfterSubscribeUrl),
	)

	if expectedUrl != requestedUrl {
		t.Fatalf(`Request URL should be "%s", "%s" given`, expectedUrl, requestedUrl)
	}

	if expectedApiKey != requestedApiKey {
		t.Fatalf(`API key should be "%s", "%s" given`, expectedApiKey, requestedApiKey)
	}

	if expectedID != requestedID {
		t.Fatalf(`ID should be %d, %d given`, expectedID, requestedID)
	}

	if expectedTitle != requestedTitle {
		t.Fatalf(`Title should be "%s", "%s" given`, expectedTitle, requestedTitle)
	}

	if expectedBeforeSubscribeUrl != requestedBeforeSubscribeUrl {
		t.Fatalf(`Param "before_subscribe_url" should be "%s", "%s" given`, expectedBeforeSubscribeUrl, requestedBeforeSubscribeUrl)
	}

	if expectedAfterSubscribeUrl != requestedAfterSubscribeUrl {
		t.Fatalf(`Param "after_subscribe_url" should be "%s", "%s" given`, expectedAfterSubscribeUrl, requestedAfterSubscribeUrl)
	}

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}
}

func TestUniSender_DeleteList(t *testing.T) {
	expectedApiKey := randomString(32)
	var requestedApiKey string

	language := randomLanguage()

	expectedUrl := fmt.Sprintf("https://api.unisender.com/%s/api/deleteList", language)
	var requestedUrl string

	var requestedID int64

	expectedID := randomInt64(999, 99999)

	client := &http.Client{
		Transport: roundTrip{
			Before: func(req *http.Request) {
				requestedUrl = req.URL.String()
				requestedApiKey = req.PostFormValue("api_key")
				requestedID, _ = strconv.ParseInt(req.PostFormValue("list_id"), 10, 64)
			},
		},
	}

	usndr := unisender.New(expectedApiKey)
	usndr.SetLanguage(language)
	usndr.SetClient(client)

	err := usndr.DeleteList(expectedID)

	if expectedUrl != requestedUrl {
		t.Fatalf(`Request URL should be "%s", "%s" given`, expectedUrl, requestedUrl)
	}

	if expectedApiKey != requestedApiKey {
		t.Fatalf(`API key should be "%s", "%s" given`, expectedApiKey, requestedApiKey)
	}

	if expectedID != requestedID {
		t.Fatalf(`ID should be %d, %d given`, expectedID, requestedID)
	}

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}
}
