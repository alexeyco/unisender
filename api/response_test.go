package api_test

import (
	"testing"

	"github.com/alexeyco/unisender/api"
)

func TestResponse_IsError(t *testing.T) {
	var response api.Response

	response = api.Response{
		Error: "Some error",
	}

	if !response.IsError() {
		t.Error("Response should have error")
	}

	response = api.Response{
		Code: "some_error_code",
	}

	if !response.IsError() {
		t.Error("Response should have error")
	}
}

var errors = map[string]error{
	"invalid_api_key":                     api.ErrInvalidAPIKey,
	"access_denied":                       api.ErrAccessDenied,
	"unknown_method":                      api.ErrUnknownMethod,
	"invalid_arg":                         api.ErrInvalidArg,
	"not_enough_money":                    api.ErrNotEnoughMoney,
	"retry_later":                         api.ErrRetryLater,
	"api_call_limit_exceeded_for_api_key": api.ErrAPICallLimitExceededForAPIKey,
	"api_call_limit_exceeded_for_ip":      api.ErrAPICallLimitExceededForIP,
}

func TestResponse_Err(t *testing.T) {
	var response api.Response

	for code, err := range errors {
		response = api.Response{
			Code: code,
		}

		if response.Err() != err {
			t.Fatalf(`Error should be "%s", "%s" given`, err.Error(), response.Err().Error())
		}
	}

	errText := "Some error"
	response = api.Response{
		Error: errText,
	}

	if response.Err().Error() != errText {
		t.Fatalf(`Error should be "%s", "%s" given`, errText, response.Err().Error())
	}
}
