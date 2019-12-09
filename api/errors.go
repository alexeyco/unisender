package api

import "errors"

// See https://www.unisender.com/en/support/api/common/api-errors/

var (
	// ErrWrongStatusCode returns when the remote server returns a status other than 200
	ErrWrongStatusCode = errors.New("remote server returns wrong status code")

	// ErrInvalidAPIKey the specified API access key is invalid. Check whether the value of api_key matches
	// the value specified in your account.
	ErrInvalidAPIKey = errors.New("invalid API key")

	// ErrAccessDenied access is denied. Check whether access to API is enabled in your personal account
	// and whether you are calling a method you do not have access rights to.
	ErrAccessDenied = errors.New("access denied")

	// ErrUnknownMethod the specified method name is incorrect.
	ErrUnknownMethod = errors.New("unknown method")

	// ErrInvalidArg an invalid value was specified for one of the method arguments.
	ErrInvalidArg = errors.New("invalid argument")

	// ErrNotEnoughMoney you do not have enough money on your account to execute the method.
	ErrNotEnoughMoney = errors.New("not enough money")

	// ErrRetryLater temporary failure. Try again later.
	ErrRetryLater = errors.New("retry later")

	// ErrAPICallLimitExceededForAPIKey the restriction on the number of calls of API methods per unit of time
	// was actuated. At the moment, it is 1200 calls per minute. For the sendEmail method, it is 60 calls.
	//
	// See https://www.unisender.com/en/support/api/common/unisender-api-limits/
	ErrAPICallLimitExceededForAPIKey = errors.New("API call limit exceeded for API key")

	// ErrAPICallLimitExceededForIP the restriction on the number of calls of API methods per unit of time
	// was actuated. At the moment, it is 1200 calls per minute. For the sendEmail method, it is 60 calls.
	//
	// See https://www.unisender.com/en/support/api/common/unisender-api-limits/
	ErrAPICallLimitExceededForIP = errors.New("API call limit exceeded for IP")
)
