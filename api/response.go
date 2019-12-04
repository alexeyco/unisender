package api

import (
	"errors"
)

// Response UniSender API response.
type Response struct {
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
	Code   string      `json:"code,omitempty"`
}

// IsError returns true if response has error.
func (r Response) IsError() bool {
	return r.Error != "" || r.Code != ""
}

// Err returns response error.
func (r Response) Err() error {
	if !r.IsError() {
		return nil
	}

	switch r.Code {
	case "invalid_api_key":
		return ErrInvalidAPIKey
	case "access_denied":
		return ErrAccessDenied
	case "unknown_method":
		return ErrUnknownMethod
	case "invalid_arg":
		return ErrInvalidArg
	case "not_enough_money":
		return ErrNotEnoughMoney
	case "retry_later":
		return ErrRetryLater
	case "api_call_limit_exceeded_for_api_key":
		return ErrAPICallLimitExceededForAPIKey
	case "api_call_limit_exceeded_for_ip":
		return ErrAPICallLimitExceededForIP
	default:
		return errors.New(r.Error)
	}
}
