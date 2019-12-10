package messages_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestCheckSMSRequest_Execute(t *testing.T) {
	expectedSMSID := test.RandomInt64(9999, 999999)
	var givenSMSID int64

	expectedResult := test.RandomString(12, 36)
	var givenResult string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSMSID, _ = strconv.ParseInt(req.FormValue("sms_id"), 10, 64)

		result := api.Response{
			Result: &messages.CheckSMSResult{
				Status: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.CheckSMS(req, expectedSMSID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSMSID != givenSMSID {
		t.Fatalf(`SMS ID should be %d, %d given`, expectedSMSID, givenSMSID)
	}

	if expectedResult != givenResult {
		t.Fatalf(`Result should be "%s", "%s" given`, expectedResult, givenResult)
	}
}
