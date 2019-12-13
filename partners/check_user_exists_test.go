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

func TestCheckUserExistsRequest_Login(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)
	var givenLogin string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLogin = req.FormValue("login")

		result := api.Response{
			Result: &partners.CheckUserExistsResult{
				LoginExists: true,
				EmailExists: false,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := partners.CheckUserExists(req).
		Login(expectedLogin).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLogin != givenLogin {
		t.Fatalf(`Login should be "%s", "%s" given`, expectedLogin, givenLogin)
	}
}

func TestCheckUserExistsRequest_Email(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	var givenEmail string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmail = req.FormValue("email")

		result := api.Response{
			Result: &partners.CheckUserExistsResult{
				LoginExists: true,
				EmailExists: false,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := partners.CheckUserExists(req).
		Email(expectedEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmail != givenEmail {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedEmail, givenEmail)
	}
}

func TestCheckUserExistsRequest_Execute(t *testing.T) {
	expectedResult := randomCheckUserExistsResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := partners.CheckUserExists(req).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	var expectedLoginExists bool
	if expectedResult["login_exists"] == 1 {
		expectedLoginExists = true
	}

	var expectedEmailExists bool
	if expectedResult["email_exists"] == 1 {
		expectedEmailExists = true
	}

	if expectedLoginExists != givenResult.LoginExists {
		t.Fatalf(`Result param "user_exists" should be "%t", "%t" given`, expectedLoginExists, givenResult.LoginExists)
	}

	if expectedEmailExists != givenResult.EmailExists {
		t.Fatalf(`Result param "email_exists" should be "%t", "%t" given`, expectedEmailExists, givenResult.EmailExists)
	}
}

func randomCheckUserExistsResult() map[string]interface{} {
	return map[string]interface{}{
		"login_exists": test.RandomInt(0, 1),
		"email_exists": test.RandomInt(0, 1),
	}
}
