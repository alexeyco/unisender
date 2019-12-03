package unisender_test

import (
	"testing"

	"github.com/alexeyco/unisender"
)

func TestResponse_IsError(t *testing.T) {
	var response unisender.Response

	response = unisender.Response{
		Error: "Some error",
	}

	if !response.IsError() {
		t.Error("Response should have error")
	}

	response = unisender.Response{
		Code: "some_error_code",
	}

	if !response.IsError() {
		t.Error("Response should have error")
	}
}

var errors = map[string]error{
	"invalid_api_key":                     unisender.ErrInvalidAPIKey,
	"access_denied":                       unisender.ErrAccessDenied,
	"unknown_method":                      unisender.ErrUnknownMethod,
	"invalid_arg":                         unisender.ErrInvalidArg,
	"not_enough_money":                    unisender.ErrNotEnoughMoney,
	"retry_later":                         unisender.ErrRetryLater,
	"api_call_limit_exceeded_for_api_key": unisender.ErrAPICallLimitExceededForAPIKey,
	"api_call_limit_exceeded_for_ip":      unisender.ErrAPICallLimitExceededForIP,
}

func TestResponse_Err(t *testing.T) {
	var response unisender.Response

	for code, err := range errors {
		response = unisender.Response{
			Code: code,
		}

		if response.Err() != err {
			t.Fatalf(`Error should be "%s", "%s" given`, err.Error(), response.Err().Error())
		}
	}

	errText := "Some error"
	response = unisender.Response{
		Error: errText,
	}

	if response.Err().Error() != errText {
		t.Fatalf(`Error should be "%s", "%s" given`, errText, response.Err().Error())
	}
}
