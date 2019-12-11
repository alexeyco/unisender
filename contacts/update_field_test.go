package contacts_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestUpdateFieldRequest_PublicName(t *testing.T) {
	expectedFieldID := test.RandomInt64(9999, 999999)
	expectedName := test.RandomString(12, 36)

	expectedPublicName := test.RandomString(12, 36)
	var givenPublicName string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenPublicName = req.FormValue("public_name")

		result := api.Response{
			Result: &contacts.UpdateFieldResult{
				ID: expectedFieldID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts.UpdateField(req, expectedFieldID, expectedName).
		PublicName(expectedPublicName).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedPublicName != givenPublicName {
		t.Fatalf(`Public name should be "%s", "%s" given`, expectedPublicName, givenPublicName)
	}
}

func TestUpdateFieldRequest_Execute(t *testing.T) {
	expectedFieldID := test.RandomInt64(9999, 999999)
	var givenFieldID int64
	var givenResult int64

	expectedName := test.RandomString(12, 36)
	var givenName string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenFieldID, _ = strconv.ParseInt(req.FormValue("id"), 10, 64)
		givenName = req.FormValue("name")

		result := api.Response{
			Result: &contacts.UpdateFieldResult{
				ID: expectedFieldID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := contacts.UpdateField(req, expectedFieldID, expectedName).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedFieldID != givenFieldID {
		t.Fatalf(`Field ID should be %d, %d given`, expectedFieldID, givenFieldID)
	}

	if expectedName != givenName {
		t.Fatalf(`Field name should be "%s", "%s" given`, expectedName, givenName)
	}

	if expectedFieldID != givenResult {
		t.Fatalf(`Result ID should be %d, %d given`, expectedFieldID, givenResult)
	}
}
