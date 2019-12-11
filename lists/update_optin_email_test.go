package lists_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/lists"
	"github.com/alexeyco/unisender/test"
)

func TestUpdateOptInEmailRequest_SenderName(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedSenderName := test.RandomString(12, 36)
	var givenSenderName string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSenderName = req.FormValue("sender_name")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := lists.UpdateOptInEmail(req, expectedListID).
		SenderName(expectedSenderName).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSenderName != givenSenderName {
		t.Fatalf(`Sender name should be "%s", "%s" given`, expectedSenderName, givenSenderName)
	}
}

func TestUpdateOptInEmailRequest_SenderEmail(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedSenderEmail := test.RandomString(12, 36)
	var givenSenderEmail string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSenderEmail = req.FormValue("sender_email")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := lists.UpdateOptInEmail(req, expectedListID).
		SenderEmail(expectedSenderEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSenderEmail != givenSenderEmail {
		t.Fatalf(`Sender email should be "%s", "%s" given`, expectedSenderEmail, givenSenderEmail)
	}
}

func TestUpdateOptInEmailRequest_Subject(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedSubject := test.RandomString(12, 36)
	var givenSubject string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSubject = req.FormValue("subject")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := lists.UpdateOptInEmail(req, expectedListID).
		Subject(expectedSubject).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSubject != givenSubject {
		t.Fatalf(`Subject should be "%s", "%s" given`, expectedSubject, givenSubject)
	}
}

func TestUpdateOptInEmailRequest_Body(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedBody := test.RandomString(12, 36)
	var givenBody string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBody = req.FormValue("body")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := lists.UpdateOptInEmail(req, expectedListID).
		Body(expectedBody).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBody != givenBody {
		t.Fatalf(`Body should be "%s", "%s" given`, expectedBody, givenBody)
	}
}

func TestUpdateOptInEmailRequest_Execute(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)
	var givenListID int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenListID, err = strconv.ParseInt(req.FormValue("list_id"), 10, 64)

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := lists.UpdateOptInEmail(req, expectedListID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedListID != givenListID {
		t.Fatalf("List ID should be %d, %d given", expectedListID, givenListID)
	}
}
