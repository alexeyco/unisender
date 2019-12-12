package partners

import "github.com/alexeyco/unisender/api"

// ValidateSenderResult email validation response.
type ValidateSenderResult struct {
	Message string `json:"message"`
}

// ValidateSenderRequest request to validate email.
type ValidateSenderRequest struct {
	request *api.Request
}

// Login sets login of the user who will be allowed to use the return address. The reseller can indicate
// the login of "its" user here. If the login is not specified, the address will be linked to the user
// calling the method.
func (r *ValidateSenderRequest) Login(login string) *ValidateSenderRequest {
	r.request.Add("login", login)
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *ValidateSenderRequest) Execute() (message string, err error) {
	var result ValidateSenderResult
	if err = r.request.Execute("validateSender", &result); err != nil {
		return
	}

	message = result.Message

	return
}

// ValidateSender sends a message to the email address with a link to confirm the address as the return address.
func ValidateSender(request *api.Request, email string) *ValidateSenderRequest {
	request.Add("email", email)

	return &ValidateSenderRequest{
		request: request,
	}
}
