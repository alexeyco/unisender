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

func TestListMessagesRequest_From(t *testing.T) {
	expectedFrom := test.RandomTime(12, 365)
	var givenFrom time.Time

	expectedResult := randomListMessagesResult()

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

	_, err := messages.ListMessages(req).
		From(expectedFrom).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedFrom != givenFrom {
		t.Fatalf(`From should be "%s", "%s" given`, expectedFrom, givenFrom)
	}
}

func TestListMessagesRequest_To(t *testing.T) {
	expectedTo := test.RandomTime(12, 365)
	var givenTo time.Time

	expectedResult := randomListMessagesResult()

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

	_, err := messages.ListMessages(req).
		To(expectedTo).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTo != givenTo {
		t.Fatalf(`To should be "%s", "%s" given`, expectedTo, givenTo)
	}
}

func TestListMessagesRequest_Limit(t *testing.T) {
	expectedLimit := test.RandomInt(9999, 999999)
	var givenLimit int

	expectedResult := randomListMessagesResult()

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

	_, err := messages.ListMessages(req).
		Limit(expectedLimit).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLimit != givenLimit {
		t.Fatalf(`Limit should be %d, %d given`, expectedLimit, givenLimit)
	}
}

func TestListMessagesRequest_Offset(t *testing.T) {
	expectedOffset := test.RandomInt(9999, 999999)
	var givenOffset int

	expectedResult := randomListMessagesResult()

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

	_, err := messages.ListMessages(req).
		Offset(expectedOffset).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedOffset != givenOffset {
		t.Fatalf(`Offset should be %d, %d given`, expectedOffset, givenOffset)
	}
}

func TestListMessagesRequest_Execute(t *testing.T) {
	expectedResult := randomListMessagesResult()
	var givenResult []messages.ListMessagesResult

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

	givenResult, err := messages.ListMessages(req).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}

func randomListMessagesResult() (res []messages.ListMessagesResult) {
	num := test.RandomInt(12, 36)
	res = make([]messages.ListMessagesResult, num)

	for i := 0; i < num; i++ {
		res[i] = messages.ListMessagesResult{
			ID:            test.RandomInt64(9999, 999999),
			SubUserLogin:  test.RandomString(12, 36),
			ListID:        test.RandomInt64(9999, 999999),
			SegmentID:     test.RandomInt64(9999, 999999),
			ServiceType:   test.RandomString(12, 36),
			Lang:          test.RandomString(12, 36),
			SenderName:    test.RandomString(12, 36),
			SenderEmail:   test.RandomString(12, 36),
			Subject:       test.RandomString(12, 36),
			MessageFormat: test.RandomString(12, 36),
		}
	}

	return
}
