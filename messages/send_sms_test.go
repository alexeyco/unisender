package messages_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/messages"
	"github.com/alexeyco/unisender/test"
)

func TestSendSMSRequest_Sender(t *testing.T) {
	expectedPhone := test.RandomStringSlice(4, 12)

	expectedSender := test.RandomString(12, 36)
	var givenSender string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenSender = req.FormValue("sender")

		result := api.Response{
			Result: randomSendSMSResult(),
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendSMS(req, expectedPhone...).
		Sender(expectedSender).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedSender != givenSender {
		t.Fatalf(`Sender should be "%s", "%s" given`, expectedSender, givenSender)
	}
}

func TestSendSMSRequest_Text(t *testing.T) {
	expectedPhone := test.RandomStringSlice(4, 12)

	expectedText := test.RandomString(12, 36)
	var givenText string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenText = req.FormValue("text")

		result := api.Response{
			Result: randomSendSMSResult(),
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := messages.SendSMS(req, expectedPhone...).
		Text(expectedText).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedText != givenText {
		t.Fatalf(`Text should be "%s", "%s" given`, expectedText, givenText)
	}
}

func TestSendSMSRequest_Execute(t *testing.T) {
	expectedPhone := test.RandomStringSlice(4, 12)
	var givenPhone []string

	expectedResult := randomSendSMSResult()
	var givenResult *messages.SendSMSResult

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenPhone = strings.Split(req.FormValue("phone"), ",")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := messages.SendSMS(req, expectedPhone...).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedPhone, givenPhone) {
		t.Fatal(`Phone numbers should be equal`)
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal(`Results should be equal`)
	}
}

func randomSendSMSResult() *messages.SendSMSResult {
	return &messages.SendSMSResult{
		Currency: test.RandomString(12, 36),
		Price:    0,
		SMSID:    test.RandomInt64(9999, 999999),
	}
}
