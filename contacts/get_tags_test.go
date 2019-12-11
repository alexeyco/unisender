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

func TestGetTagsRequest_Execute(t *testing.T) {
	expectedResult := randomGetTagsResult()

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

	givenResult, err := contacts.GetTags(req).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal(`Results should be equal`)
	}
}

func randomGetTagsResult() []contacts.GetTagsResult {
	n := test.RandomInt(12, 36)
	res := make([]contacts.GetTagsResult, n)

	for i := 0; i < n; i++ {
		res[i] = contacts.GetTagsResult{
			ID:   test.RandomInt64(9999, 999999),
			Name: test.RandomString(12, 36),
		}
	}

	return res
}
