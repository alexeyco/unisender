package contacts_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestDeleteTagRequest_Execute(t *testing.T) {
	expectedTagID := test.RandomInt64(9999, 999999)
	var givenTagID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTagID, _ = strconv.ParseInt(req.FormValue("id"), 10, 64)

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := contacts.DeleteTag(req, expectedTagID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTagID != givenTagID {
		t.Fatalf(`Tag ID should be %d, %d given`, expectedTagID, givenTagID)
	}
}
