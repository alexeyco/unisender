package partners_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/partners"
	"github.com/alexeyco/unisender/test"
)

func TestValidateSenderRequest_Login(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)

	expectedLogin := test.RandomString(12, 36)
	var givenLogin string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLogin = req.FormValue("login")

		result := api.Response{
			Result: &partners.ValidateSenderResult{
				Message: test.RandomString(12, 36),
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := partners.ValidateSender(req, expectedEmail).
		Login(expectedLogin).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLogin != givenLogin {
		t.Fatalf(`Login should be "%s", "%s" given`, expectedLogin, givenLogin)
	}
}

func TestValidateSenderRequest_Execute(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	var givenEmail string

	expectedMessage := test.RandomString(12, 36)
	var givenMessage string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmail = req.FormValue("email")

		result := api.Response{
			Result: &partners.ValidateSenderResult{
				Message: expectedMessage,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenMessage, err := partners.ValidateSender(req, expectedEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmail != givenEmail {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedEmail, givenEmail)
	}

	if expectedMessage != givenMessage {
		t.Fatalf(`Message should be "%s", "%s" given`, expectedMessage, givenMessage)
	}
}
