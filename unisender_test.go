package unisender

import (
	"net/http"
	"testing"

	"github.com/alexeyco/unisender/test"
)

func TestUniSender_ApiKey(t *testing.T) {
	apiKeyExpected := "foo-bar-api-key"
	var apiKeyRequested string

	c := test.NewClient(func(req *http.Request) (res *http.Response, err error) {
		apiKeyRequested = req.FormValue("api_key")
		res = &http.Response{
			StatusCode: http.StatusOK,
		}

		return
	})

	usndr := New(apiKeyExpected)
	usndr.SetClient(c)

	err := usndr.DeleteList(123).Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if apiKeyExpected != apiKeyRequested {
		t.Fatalf(`API key should be "%s", "%s" given`, apiKeyExpected, apiKeyRequested)
	}
}
