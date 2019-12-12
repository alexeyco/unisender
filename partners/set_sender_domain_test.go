package partners_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/partners"
	"github.com/alexeyco/unisender/test"
)

func TestSetSenderDomainRequest_Execute(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)
	var givenLogin string

	expectedDomain := test.RandomString(12, 36)
	var givenDomain string

	expectedResult := &partners.SetSenderDomainResult{
		DKIM: test.RandomString(12, 36),
	}

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLogin = req.FormValue("username")
		givenDomain = req.FormValue("domain")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := partners.SetSenderDomain(req, expectedLogin, expectedDomain).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLogin != givenLogin {
		t.Fatalf(`Login should be "%s", "%s" given`, expectedLogin, givenLogin)
	}

	if expectedDomain != givenDomain {
		t.Fatalf(`Domain should be "%s", "%s" given`, expectedDomain, givenDomain)
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatalf("Results should be equal")
	}
}
