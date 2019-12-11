package messages_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestGetActualMessageVersionRequest_Execute(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	var givenMessageID int64

	expectedResult := &messages.GetActualMessageVersionResult{
		MessageID:       test.RandomInt64(9999, 999999),
		ActualVersionID: test.RandomInt64(9999, 999999),
	}

	var givenResult *messages.GetActualMessageVersionResult

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageID, _ = strconv.ParseInt(req.FormValue("message_id"), 10, 64)

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.GetActualMessageVersion(req, expectedMessageID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedMessageID != givenMessageID {
		t.Fatalf(`Message ID should be %d, %d given`, expectedMessageID, givenMessageID)
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}
