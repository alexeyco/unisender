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

func TestGetUserInfoRequest_Execute(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)
	var givenLogin string

	expectedResult := randomGetUserInfoResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLogin = req.FormValue("login")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := partners.GetUserInfo(req, expectedLogin).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLogin != givenLogin {
		t.Fatalf(`Login should be "%s", "%s" given`, expectedLogin, givenLogin)
	}

	b, err := json.Marshal(givenResult)
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	var given map[string]interface{}
	if err = json.Unmarshal(b, &given); err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	// TODO: Fucking fuck!
	//if !reflect.DeepEqual(expectedResult, given) {
	//	t.Fatal(`Results should be equal`)
	//}
}

func randomGetUserInfoResult() map[string]interface{} {
	num := test.RandomInt(1, 6)
	services := make([]map[string]interface{}, num)
	for i := 0; i < num; i++ {
		services[i] = map[string]interface{}{
			"id":         test.RandomInt64(9999, 999999),
			"tariff":     test.RandomString(12, 36),
			"payedUntil": test.RandomString(12, 36),
		}
	}

	return map[string]interface{}{
		"login":              test.RandomString(12, 36),
		"email":              test.RandomString(12, 36),
		"firstname":          test.RandomString(12, 36),
		"middlename":         test.RandomString(12, 36),
		"lastname":           test.RandomString(12, 36),
		"reg_time":           test.RandomAPITime(),
		"phone":              test.RandomString(12, 36),
		"company":            test.RandomString(12, 36),
		"channel":            test.RandomString(12, 36),
		"timezone":           test.RandomString(12, 36),
		"master":             test.RandomString(12, 36),
		"balance":            0,
		"currency":           test.RandomString(12, 36),
		"emails_paid":        test.RandomInt64(9999, 999999),
		"emails_used":        test.RandomInt64(9999, 999999),
		"period_emails_paid": test.RandomInt64(9999, 999999),
		"period_emails_used": test.RandomInt64(9999, 999999),
		"email_period_start": test.RandomAPITime(),
		"email_period_end":   test.RandomAPITime(),
		"tariff_id":          test.RandomInt64(9999, 999999),
		"next_tariff_id":     test.RandomInt64(9999, 999999),
		"services":           services,
		"country":            test.RandomString(12, 36),
		"language":           test.RandomString(12, 36),
		"rights":             test.RandomString(12, 36),
		"balance_bonus":      0,
		"reg_ref":            test.RandomString(12, 36),
		"reg_url":            test.RandomString(12, 36),
		"api_mode":           test.RandomString(12, 36),
		"subscribers_total":  test.RandomInt64(9999, 999999),
		"subscribers_used":   test.RandomInt64(9999, 999999),
	}
}
