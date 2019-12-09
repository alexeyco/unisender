package messages

import "github.com/alexeyco/unisender/api"

// ValidateSenderResponse email validation response.
type ValidateSenderResponse struct {
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
	var response ValidateSenderResponse
	if err = r.request.Execute("validateSender", &response); err != nil {
		return
	}

	message = response.Message

	return
}

// ValidateSender sends a message to the email address with a link to confirm the address as the return address.
// After clicking on this link, you can send messages on behalf of this email address.
//
// See: https://www.unisender.com/en/support/api/messages/validatesender/
func ValidateSender(request *api.Request, email string) *ValidateSenderRequest {
	request.Add("email", email)

	return &ValidateSenderRequest{
		request: request,
	}
}
