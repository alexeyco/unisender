package contacts_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestDeleteListRequest_Execute(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)
	var givenListID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenListID, _ = strconv.ParseInt(req.FormValue("list_id"), 10, 64)

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := contacts.DeleteList(req, expectedListID).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedListID != givenListID {
		t.Fatalf("List ID should be %d, %d given", expectedListID, givenListID)
	}
}
