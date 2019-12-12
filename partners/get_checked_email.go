package partners

import "github.com/alexeyco/unisender/api"

// GetCheckedEmailResponse sender info.
type GetCheckedEmailResponse struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	IsChecked bool   `json:"is_checked"`
}

// GetCheckedEmailRequest request, that gets an object with confirmed and unconfirmed sender’s addresses.
type GetCheckedEmailRequest struct {
	request *api.Request
}

// Email confirmation status at the specified address. If skipped, the method will return statuses
// for all confirmed/unconfirmed addresses.
func (r *GetCheckedEmailRequest) Email(email string) *GetCheckedEmailRequest {
	r.request.Add("email", email)
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *GetCheckedEmailRequest) Execute() (res []GetCheckedEmailResponse, err error) {
	err = r.request.Execute("getCheckedEmail", &res)
	return
}

// GetCheckedEmail returns request, that gets an object with confirmed and unconfirmed sender’s addresses.
func GetCheckedEmail(request *api.Request, login string) *GetCheckedEmailRequest {
	request.Add("login", login)

	return &GetCheckedEmailRequest{
		request: request,
	}
}
