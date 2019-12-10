package contacts_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestDeleteFieldRequest_Execute(t *testing.T) {
	expectedFieldID := test.RandomInt64(9999, 999999)
	var givenFieldID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenFieldID, _ = strconv.ParseInt(req.FormValue("id"), 10, 64)

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := contacts.DeleteField(req, expectedFieldID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedFieldID != givenFieldID {
		t.Fatalf(`Field ID should be %d, %d given`, expectedFieldID, givenFieldID)
	}
}
