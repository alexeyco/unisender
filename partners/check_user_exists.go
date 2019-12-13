package partners

import "github.com/alexeyco/unisender/api"

// CheckUserExistsResult request to check user exists result.
type CheckUserExistsResult struct {
	LoginExists bool
	EmailExists bool
}

type checkUserExistsResultRaw struct {
	LoginExists int `json:"login_exists"`
	EmailExists int `json:"email_exists"`
}

// CheckUserExistsRequest
type CheckUserExistsRequest struct {
	request *api.Request
}

// Login sets login to check.
func (r *CheckUserExistsRequest) Login(login string) *CheckUserExistsRequest {
	r.request.Add("login", login)
	return r
}

// Email sets email to check.
func (r *CheckUserExistsRequest) Email(email string) *CheckUserExistsRequest {
	r.request.Add("email", email)
	return r
}

// Execute sends request to UniSender API and returns result.
func (r *CheckUserExistsRequest) Execute() (res *CheckUserExistsResult, err error) {
	var result checkUserExistsResultRaw
	if err = r.request.Execute("checkUserExists", &result); err != nil {
		return
	}

	res = &CheckUserExistsResult{}
	if result.LoginExists == 1 {
		res.LoginExists = true
	}

	if result.EmailExists == 1 {
		res.EmailExists = true
	}

	return
}

// CheckUserExists returns request to check whether the user with the given login and email has been registered. It is intended
// to be called only by resellers.
func CheckUserExists(request *api.Request) *CheckUserExistsRequest {
	return &CheckUserExistsRequest{
		request: request,
	}
}
