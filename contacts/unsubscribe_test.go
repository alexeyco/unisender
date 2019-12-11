package contacts_test

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestUnsubscribeRequest_ContactTypeEmail(t *testing.T) {
	expectedContact := test.RandomString(12, 36)

	expectedContactType := "email"
	var givenContactType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContactType = req.FormValue("contact_type")

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

	if expectedContactType != givenContactType {
		t.Fatalf(`Contact type should be "%s", "%s" given`, expectedContactType, givenContactType)
	}
}

func TestUnsubscribeRequest_ContactTypePhone(t *testing.T) {
	expectedContact := test.RandomString(12, 36)

	expectedContactType := "phone"
	var givenContactType string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContactType = req.FormValue("contact_type")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := contacts.Unsubscribe(req, expectedContact).
		ContactTypePhone().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedContactType != givenContactType {
		t.Fatalf(`Contact type should be "%s", "%s" given`, expectedContactType, givenContactType)
	}
}

func TestUnsubscribeRequest_ListIDs(t *testing.T) {
	expectedContact := test.RandomString(12, 36)

	expectedListIDs := test.RandomInt64Slice(12, 36)
	var givenListIDs []int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		s := req.FormValue("list_ids")
		for _, idStr := range strings.Split(s, ",") {
			id, _ := strconv.ParseInt(idStr, 10, 64)
			givenListIDs = append(givenListIDs, id)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := contacts.Unsubscribe(req, expectedContact).
		ListIDs(expectedListIDs...).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedListIDs, givenListIDs) {
		t.Fatal("List IDs should be equal")
	}
}

func TestUnsubscribeRequest_Execute(t *testing.T) {
	expectedContact := test.RandomString(12, 36)
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
