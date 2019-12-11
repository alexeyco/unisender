package contacts_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestGetTotalContactCountRequest_Execute(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)
	var givenLogin string

	expectedResult := test.RandomInt64(9999, 999999)
	var givenResult int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLogin = req.FormValue("login")

		result := api.Response{
			Result: &contacts.GetTotalContactsCountResult{
				Total: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := contacts.GetTotalContactsCount(req, expectedLogin).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLogin != givenLogin {
		t.Fatalf(`Login should be "%s", "%s" given`, expectedLogin, givenLogin)
	}

	if expectedResult != givenResult {
		t.Fatalf(`Total contacts count should be %d, %d given`, expectedResult, givenResult)
	}
}
