package contacts_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts"
	"github.com/alexeyco/unisender/test"
)

func TestGetContactFieldValuesRequest_Execute(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	var givenEmail string

	expectedFieldIDs := test.RandomInt64Slice(12, 36)
	var givenFieldIDs []int64

	expectedResult := randomFieldValues()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmail = req.FormValue("email")
		fieldIDs := strings.Split(req.FormValue("field_ids"), ",")
		for _, fieldID := range fieldIDs {
			id, _ := strconv.ParseInt(fieldID, 10, 64)
			givenFieldIDs = append(givenFieldIDs, id)
		}

		result := api.Response{
			Result: &contacts.GetContactFieldValuesResult{
				FieldValues: expectedResult,
			},
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := contacts.GetContactFieldValues(req, expectedEmail, expectedFieldIDs...).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmail != givenEmail {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedEmail, givenEmail)
	}

	if !reflect.DeepEqual(expectedFieldIDs, givenFieldIDs) {
		t.Fatal(`Field IDs should be equal`)
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal(`Results should be equal`)
	}
}

func randomFieldValues() map[int64]string {
	n := test.RandomInt(12, 32)
	fieldValues := make(map[int64]string, n)

	for i := 0; i < n; i++ {
		fieldValues[test.RandomInt64(9999, 999999)] = test.RandomString(12, 36)
	}

	return fieldValues
}
