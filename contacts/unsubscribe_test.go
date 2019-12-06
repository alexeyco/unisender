package contacts_test

import (
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestUnsubscribeRequest_ContactTypeEmail(t *testing.T) {

}

func TestUnsubscribeRequest_ContactTypePhone(t *testing.T) {

}

func TestUnsubscribeRequest_ListIDs(t *testing.T) {

}

func TestUnsubscribeRequest_Execute(t *testing.T) {
	expectedContact := "foo@bar.example"
	var givenContact string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContact = req.FormValue("contact")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := contacts.Unsubscribe(req, expectedContact).
		ContactTypeEmail().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedContact != givenContact {
		t.Fatalf(`Contact should be "%s", "%s" given`, expectedContact, givenContact)
	}
}
