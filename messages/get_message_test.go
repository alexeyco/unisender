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

func TestGetMessageRequest_Execute(t *testing.T) {
	expectedMessageID := test.RandomInt64(9999, 999999)
	var givenMessageID int64

	expectedResult := &messages.GetMessageResult{
		ID:              test.RandomInt64(9999, 999999),
		SubUserLogin:    test.RandomString(12, 36),
		SenderEmail:     test.RandomString(12, 36),
		SenderName:      test.RandomString(12, 36),
		Sender:          test.RandomString(12, 36),
		Subject:         test.RandomString(12, 36),
		Body:            test.RandomString(12, 36),
		BodyText:        test.RandomString(12, 36),
		ListID:          test.RandomInt64(9999, 999999),
		LastUpdate:      test.RandomTime(12, 365),
		ServiceType:     test.RandomString(12, 36),
		Lang:            test.RandomString(12, 36),
		ActualVersionID: test.RandomInt64(9999, 999999),
		MessageFormat:   test.RandomString(12, 36),
		WrapType:        test.RandomString(12, 36),
		ImagesBehavior:  test.RandomString(12, 36),
	}

	var givenResult *messages.GetMessageResult

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenMessageID, _ = strconv.ParseInt(req.FormValue("id"), 10, 64)

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.GetMessage(req, expectedMessageID).
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
