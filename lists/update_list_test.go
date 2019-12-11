package lists_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/lists"
	"github.com/alexeyco/unisender/test"
)

func TestUpdateListRequest_BeforeSubscribeUrl(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)
	expectedTitle := test.RandomString(12, 36)

	expectedBeforeSubscribeUrl := test.RandomString(12, 36)
	var givenBeforeSubscribeUrl string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBeforeSubscribeUrl = req.FormValue("before_subscribe_url")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := lists.UpdateList(req, expectedListID, expectedTitle).
		BeforeSubscribeUrl(expectedBeforeSubscribeUrl).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBeforeSubscribeUrl != givenBeforeSubscribeUrl {
		t.Fatalf(`Before subscribe URL should be "%s", "%s" given`, expectedBeforeSubscribeUrl, givenBeforeSubscribeUrl)
	}
}

func TestUpdateListRequest_AfterSubscribeUrl(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)
	expectedTitle := test.RandomString(12, 36)

	expectedAfterSubscribeUrl := test.RandomString(12, 36)
	var givenAfterSubscribeUrl string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenAfterSubscribeUrl = req.FormValue("after_subscribe_url")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := lists.UpdateList(req, expectedListID, expectedTitle).
		AfterSubscribeUrl(expectedAfterSubscribeUrl).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedAfterSubscribeUrl != givenAfterSubscribeUrl {
		t.Fatalf(`After subscribe URL should be "%s", "%s" given`, expectedAfterSubscribeUrl, givenAfterSubscribeUrl)
	}
}

func TestUpdateListRequest_Execute(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)
	var givenListID int64

	expectedTitle := test.RandomString(12, 36)
	var givenTitle string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenListID, err = strconv.ParseInt(req.FormValue("list_id"), 10, 64)
		givenTitle = req.FormValue("title")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := lists.UpdateList(req, expectedListID, expectedTitle).Execute()

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
