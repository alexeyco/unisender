package unisender

import "errors"

var (
	ErrWrongStatusCode               = errors.New("remote server returned wrong status code")
	ErrInvalidAPIKey                 = errors.New("invalid API key")
	ErrAccessDenied                  = errors.New("access denied")
	ErrUnknownMethod                 = errors.New("unknown method")
	ErrInvalidArg                    = errors.New("invalid argument")
	ErrNotEnoughMoney                = errors.New("not enough money")
	ErrRetryLater                    = errors.New("retry later")
	ErrAPICallLimitExceededForAPIKey = errors.New("API call limit exceeded for API key")
	ErrAPICallLimitExceededForIP     = errors.New("API call limit exceeded for IP")
)
