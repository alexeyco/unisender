package contacts_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestGetFieldsRequest_Execute(t *testing.T) {
	expectedResult := randomGetFieldsResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := contacts.GetFields(req).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal(`Results should be equal`)
	}
}

func randomGetFieldsResult() []contacts.GetFieldsResult {
	n := test.RandomInt(12, 36)
	res := make([]contacts.GetFieldsResult, n)

	for i := 0; i < n; i++ {
		res[i] = contacts.GetFieldsResult{
			ID:        test.RandomInt64(9999, 999999),
			Name:      test.RandomString(12, 36),
			IsVisible: test.RandomInt(0, 1),
			ViewPos:   test.RandomInt(0, 999),
		}
	}

	return res
}
