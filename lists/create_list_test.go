package lists_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/lists"
)

func TestCreateListRequest_Execute(t *testing.T) {
	expectedListID := int64(randomInt(9999, 999999))

	expectedTitle := fmt.Sprintf("Title #%d", randomInt(9999, 999999))
	var givenTitle string

	expectedBeforeSubscribeUrl := "https://before-subscribe.url"
	var givenBeforeSubscribeUrl string

	expectedAfterSubscribeUrl := "https://after-subscribe.url"
	var givenAfterSubscribeUrl string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTitle = req.FormValue("title")
		givenBeforeSubscribeUrl = req.FormValue("before_subscribe_url")
		givenAfterSubscribeUrl = req.FormValue("after_subscribe_url")

		response := fmt.Sprintf(`{"result":{"id":%d}}`, expectedListID)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
		}, nil
	})

	givenListsID, err := lists.CreateList(req, expectedTitle).
		BeforeSubscribeUrl(expectedBeforeSubscribeUrl).
		AfterSubscribeUrl(expectedAfterSubscribeUrl).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedListID != givenListsID {
		t.Fatalf("List ID should be %d, %d given", expectedListID, givenListsID)
	}

	if expectedTitle != givenTitle {
		t.Fatalf(`Title should be "%s", "%s" given`, expectedTitle, givenTitle)
	}

	if givenBeforeSubscribeUrl != expectedBeforeSubscribeUrl {
		t.Errorf(`Param "before_subscribe_url" should be "%s", "%s" given`, expectedBeforeSubscribeUrl, givenBeforeSubscribeUrl)
	}

	if givenAfterSubscribeUrl != expectedAfterSubscribeUrl {
		t.Errorf(`Param "after_subscribe_url" should be "%s", "%s" given`, expectedAfterSubscribeUrl, givenAfterSubscribeUrl)
	}
}
