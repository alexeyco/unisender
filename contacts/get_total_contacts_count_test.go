package contacts_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts"
)

func TestGetTotalContactCountRequest_Execute(t *testing.T) {
	expectedLogin := randomString(12, 36)
	var givenLogin string

	expectedResult := randomGetTotalContactsCountResult()

	req := newRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLogin = req.FormValue("login")

		result := api.Response{
			Result: expectedResult,
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

	if expectedResult.Total != givenResult {
		t.Fatalf(`Total count should be %d, %d given`, expectedResult.Total, givenResult)
	}
}

func randomGetTotalContactsCountResult() *contacts.GetTotalContactsCountResponse {
	return &contacts.GetTotalContactsCountResponse{
		Total: int64(randomInt(9999, 999999)),
	}
}
