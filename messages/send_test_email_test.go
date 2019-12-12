package messages_test

import (
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
	"net/http"
	"strconv"
	"testing"
)

func TestSendTestEmailRequest_To(t *testing.T) {
	expectedEmailID := test.RandomInt64(9999, 999999)

	expectedTo := test.RandomString(12, 36)
	var givenTo string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTo = req.FormValue("email")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.SendTestEmail(req, expectedEmailID).
		To(expectedTo).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTo != givenTo {
		t.Fatalf(`To should be "%s", "%s" given`, expectedTo, givenTo)
	}
}

func TestSendTestEmailRequest_Execute(t *testing.T) {
	expectedEmailID := test.RandomInt64(9999, 999999)
	var givenEmailID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmailID, _ = strconv.ParseInt(req.FormValue("id"), 10, 64)

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := messages.SendTestEmail(req, expectedEmailID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmailID != givenEmailID {
		t.Fatalf(`Email ID should be %d, %d given`, expectedEmailID, givenEmailID)
	}
}
