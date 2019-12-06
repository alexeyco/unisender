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

func TestGetListsRequest_Execute(t *testing.T) {
	expectedLists := getRandomListsSlice()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		result := api.Response{
			Result: expectedLists,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenLists, err := contacts.GetLists(req).Execute()
	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if len(givenLists) != len(expectedLists) {
		t.Fatalf(`Lists slice should have length %d, %d given`, len(expectedLists), len(givenLists))
	}

	if !reflect.DeepEqual(expectedLists, givenLists) {
		t.Fatal("Expected and given lists should be equal")
	}
}

func getRandomListsSlice() (slice []contacts.List) {
	l := test.RandomInt(12, 36)
	for i := 0; i < l; i++ {
		slice = append(slice, contacts.List{
			ID:    test.RandomInt64(9999, 999999),
			Title: test.RandomString(12, 36),
		})
	}

	return
}
