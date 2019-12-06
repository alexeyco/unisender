package contacts_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

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

	err := contacts.UpdateList(req, expectedListID, expectedTitle).Execute()

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
