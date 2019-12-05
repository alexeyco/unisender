package contacts_test

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/alexeyco/unisender/contacts"
)

func TestExcludeRequest_ContactTypeEmail(t *testing.T) {
	expectedContact := "foo@bar.example"

	expectedContactType := "email"
	var givenContactType string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContactType = req.FormValue("contact_type")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := contacts.Exclude(req, expectedContact).
		ContactTypeEmail().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedContactType != givenContactType {
		t.Fatalf(`Contact type should be "%s", "%s" given`, expectedContactType, givenContactType)
	}
}

func TestExcludeRequest_ContactTypePhone(t *testing.T) {
	expectedContact := "+1234567890"

	expectedContactType := "phone"
	var givenContactType string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContactType = req.FormValue("contact_type")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := contacts.Exclude(req, expectedContact).
		ContactTypePhone().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedContactType != givenContactType {
		t.Fatalf(`Contact type should be "%s", "%s" given`, expectedContactType, givenContactType)
	}
}

func TestExcludeRequest_ListIDs(t *testing.T) {
	expectedContact := "foo@bar.example"

	expectedListIDs := randomInt64Slice(12, 36)
	var givenListIDs []int64

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		s := req.FormValue("list_ids")
		for _, idStr := range strings.Split(s, ",") {
			id, _ := strconv.ParseInt(idStr, 10, 64)
			givenListIDs = append(givenListIDs, id)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := contacts.Exclude(req, expectedContact).
		ContactTypeEmail().
		ListIDs(expectedListIDs...).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedListIDs, givenListIDs) {
		t.Fatal("List IDs should be equal")
	}
}

func TestExcludeRequest_Execute(t *testing.T) {
	expectedContact := "foo@bar.example"
	var givenContact string

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenContact = req.FormValue("contact")

		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	})

	err := contacts.Exclude(req, expectedContact).
		ContactTypeEmail().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedContact != givenContact {
		t.Fatalf(`Contact should be "%s", "%s" given`, expectedContact, givenContact)
	}
}

func randomInt64Slice(min, max int) (slice []int64) {
	l := randomInt(min, max)
	for i := 0; i < l; i++ {
		slice = append(slice, int64(randomInt(9999, 999999)))
	}

	return slice
}
