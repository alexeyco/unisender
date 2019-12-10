package contacts2_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts2"
	"github.com/alexeyco/unisender/test"
)

func TestGetTotalContactCountRequest_Execute(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)
	var givenLogin string

	expectedResult := randomGetTotalContactsCountResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
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

	givenResult, err := contacts2.GetTotalContactsCount(req, expectedLogin).
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

func randomGetTotalContactsCountResult() *contacts2.GetTotalContactsCountResponse {
	return &contacts2.GetTotalContactsCountResponse{
		Total: test.RandomInt64(9999, 999999),
	}
}
