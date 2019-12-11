package messages_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestCreateSMSMessageRequest_Body(t *testing.T) {
	expectedSender := test.RandomString(9999, 999999)

	expectedBody := test.RandomString(12, 36)
	var givenBody string

	expectedResult := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBody = req.FormValue("body")

		result := api.Response{
			Result: &messages.CreateSMSMessageResult{
				MessageID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateSMSMessage(req, expectedSender).
		Body(expectedBody).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedBody != givenBody {
		t.Fatalf(`Body should be "%s", "%s" given`, expectedBody, givenBody)
	}
}

func TestCreateSMSMessageRequest_ListID(t *testing.T) {
	expectedSender := test.RandomString(9999, 999999)

	expectedListID := test.RandomInt64(9999, 999999)
	var givenListID int64

	expectedResult := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenListID, _ = strconv.ParseInt(req.FormValue("list_id"), 10, 64)

		result := api.Response{
			Result: &messages.CreateSMSMessageResult{
				MessageID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateSMSMessage(req, expectedSender).
		ListID(expectedListID).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedListID != givenListID {
		t.Fatalf(`List ID should be %d, %d given`, expectedListID, givenListID)
	}
}

func TestCreateSMSMessageRequest_Tag(t *testing.T) {
	expectedSender := test.RandomString(9999, 999999)

	expectedTag := test.RandomString(12, 36)
	var givenTag string

	expectedResult := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTag = req.FormValue("tag")

		result := api.Response{
			Result: &messages.CreateSMSMessageResult{
				MessageID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateSMSMessage(req, expectedSender).
		Tag(expectedTag).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedTag != givenTag {
		t.Fatalf(`Body should be "%s", "%s" given`, expectedTag, givenTag)
	}
}

func TestCreateSMSMessageRequest_Categories(t *testing.T) {
	expectedSender := test.RandomString(9999, 999999)

	expectedCategories := test.RandomStringSlice(12, 36)
	var givenCategories []string

	expectedResult := test.RandomInt64(9999, 999999)

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenCategories = strings.Split(req.FormValue("categories"), ",")

		result := api.Response{
			Result: &messages.CreateSMSMessageResult{
				MessageID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.CreateSMSMessage(req, expectedSender).
		Categories(expectedCategories...).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedCategories, givenCategories) {
		t.Fatal(`Categories should be equal`)
	}
}

func TestCreateSMSMessageRequest_Execute(t *testing.T) {
	expectedSender := test.RandomString(9999, 999999)
	var givenSender string

	expectedResult := test.RandomInt64(9999, 999999)
	var givenResult int64

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSender = req.FormValue("sender")

		result := api.Response{
			Result: &messages.CreateSMSMessageResult{
				MessageID: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.CreateSMSMessage(req, expectedSender).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSender != givenSender {
		t.Fatalf(`List ID should be "%s", "%s" given`, expectedSender, givenSender)
	}

	if expectedResult != givenResult {
		t.Fatalf("Message ID should be %d, %d given", expectedResult, givenResult)
	}
}
