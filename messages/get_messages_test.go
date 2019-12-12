package messages_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestGetMessagesRequest_From(t *testing.T) {
	expectedFrom := test.RandomTime(12, 365)
	var givenFrom time.Time

	expectedResult := randomGetMessagesResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenFrom, _ = time.Parse("2006-01-02 15:04:05", req.FormValue("date_from"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.GetMessages(req).
		From(expectedFrom).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedFrom != givenFrom {
		t.Fatalf(`From should be "%s", "%s" given`, expectedFrom, givenFrom)
	}
}

func TestGetMessagesRequest_To(t *testing.T) {
	expectedTo := test.RandomTime(12, 365)
	var givenTo time.Time

	expectedResult := randomGetMessagesResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTo, _ = time.Parse("2006-01-02 15:04:05", req.FormValue("date_to"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.GetMessages(req).
		To(expectedTo).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTo != givenTo {
		t.Fatalf(`To should be "%s", "%s" given`, expectedTo, givenTo)
	}
}

func TestGetMessagesRequest_Limit(t *testing.T) {
	expectedLimit := test.RandomInt(9999, 999999)
	var givenLimit int

	expectedResult := randomGetMessagesResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLimit, _ = strconv.Atoi(req.FormValue("limit"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.GetMessages(req).
		Limit(expectedLimit).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLimit != givenLimit {
		t.Fatalf(`Limit should be %d, %d given`, expectedLimit, givenLimit)
	}
}

func TestGetMessagesRequest_Offset(t *testing.T) {
	expectedOffset := test.RandomInt(9999, 999999)
	var givenOffset int

	expectedResult := randomGetMessagesResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenOffset, _ = strconv.Atoi(req.FormValue("offset"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.GetMessages(req).
		Offset(expectedOffset).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedOffset != givenOffset {
		t.Fatalf(`Offset should be %d, %d given`, expectedOffset, givenOffset)
	}
}

func TestGetMessagesRequest_Execute(t *testing.T) {
	expectedResult := randomGetMessagesResult()
	var givenResult []messages.GetMessagesResult

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.GetMessages(req).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}

func randomGetMessagesResult() (res []messages.GetMessagesResult) {
	num := test.RandomInt(12, 36)
	res = make([]messages.GetMessagesResult, num)

	for i := 0; i < num; i++ {
		res[i] = messages.GetMessagesResult{
			ID:              test.RandomInt64(9999, 999999),
			SubUserLogin:    test.RandomString(12, 36),
			ListID:          test.RandomInt64(9999, 999999),
			SegmentID:       test.RandomInt64(9999, 999999),
			ServiceType:     test.RandomString(12, 36),
			ActualVersionID: test.RandomInt64(9999, 999999),
			Lang:            test.RandomString(12, 36),
			SenderName:      test.RandomString(12, 36),
			SenderEmail:     test.RandomString(12, 36),
			SMSFrom:         test.RandomString(12, 36),
			Subject:         test.RandomString(12, 36),
			Body:            test.RandomString(12, 36),
			MessageFormat:   test.RandomString(12, 36),
		}
	}

	return
}
