package contacts_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/contacts"
)

func TestCreateListRequest_BeforeSubscribeUrl(t *testing.T) {
	expectedListID := int64(randomInt(9999, 999999))
	expectedTitle := fmt.Sprintf("Title #%d", randomInt(9999, 999999))

	expectedBeforeSubscribeUrl := "https://before-subscribe.url"
	var givenBeforeSubscribeUrl string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBeforeSubscribeUrl = req.FormValue("before_subscribe_url")

		response := fmt.Sprintf(`{"result":{"id":%d}}`, expectedListID)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.CreateList(req, expectedTitle).
		BeforeSubscribeUrl(expectedBeforeSubscribeUrl).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if givenBeforeSubscribeUrl != expectedBeforeSubscribeUrl {
		t.Errorf(`Param "before_subscribe_url" should be "%s", "%s" given`, expectedBeforeSubscribeUrl, givenBeforeSubscribeUrl)
	}
}

func TestCreateListRequest_AfterSubscribeUrl(t *testing.T) {
	expectedListID := int64(randomInt(9999, 999999))
	expectedTitle := fmt.Sprintf("Title #%d", randomInt(9999, 999999))

	expectedAfterSubscribeUrl := "https://after-subscribe.url"
	var givenAfterSubscribeUrl string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenAfterSubscribeUrl = req.FormValue("after_subscribe_url")

		response := fmt.Sprintf(`{"result":{"id":%d}}`, expectedListID)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	_, err := contacts.CreateList(req, expectedTitle).
		AfterSubscribeUrl(expectedAfterSubscribeUrl).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if givenAfterSubscribeUrl != expectedAfterSubscribeUrl {
		t.Errorf(`Param "after_subscribe_url" should be "%s", "%s" given`, expectedAfterSubscribeUrl, givenAfterSubscribeUrl)
	}
}

func TestCreateListRequest_Execute(t *testing.T) {
	expectedListID := int64(randomInt(9999, 999999))

	expectedTitle := fmt.Sprintf("Title #%d", randomInt(9999, 999999))
	var givenTitle string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTitle = req.FormValue("title")

		response := fmt.Sprintf(`{"result":{"id":%d}}`, expectedListID)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	givenListID, err := contacts.CreateList(req, expectedTitle).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedListID != givenListID {
		t.Fatalf("List ID should be %d, %d given", expectedListID, givenListID)
	}

	if expectedTitle != givenTitle {
		t.Fatalf(`Title should be "%s", "%s" given`, expectedTitle, givenTitle)
	}
}
