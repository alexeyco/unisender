package contacts2_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/alexeyco/unisender/api"
	"github.com/alexeyco/unisender/contacts2"
	"github.com/alexeyco/unisender/test"
)

func TestGetContactRequest_IncludeLists(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	expectedIncludeLists := 1
	var givenIncludeLists int

	expectedResult := randomPerson()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenIncludeLists, _ = strconv.Atoi(req.FormValue("include_lists"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts2.GetContact(req, expectedEmail).
		IncludeLists().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedIncludeLists != givenIncludeLists {
		t.Fatalf(`Param "include_lists" should be %d, %d given`, expectedIncludeLists, givenIncludeLists)
	}
}

func TestGetContactRequest_IncludeFields(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	expectedIncludeFields := 1
	var givenIncludeFields int

	expectedResult := randomPerson()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenIncludeFields, _ = strconv.Atoi(req.FormValue("include_fields"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts2.GetContact(req, expectedEmail).
		IncludeFields().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedIncludeFields != givenIncludeFields {
		t.Fatalf(`Param "include_fields" should be %d, %d given`, expectedIncludeFields, givenIncludeFields)
	}
}

func TestGetContactRequest_IncludeDetails(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	expectedIncludeDetails := 1
	var givenIncludeDetails int

	expectedResult := randomPerson()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenIncludeDetails, _ = strconv.Atoi(req.FormValue("include_details"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := contacts2.GetContact(req, expectedEmail).
		IncludeDetails().
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedIncludeDetails != givenIncludeDetails {
		t.Fatalf(`Param "include_details" should be %d, %d given`, expectedIncludeDetails, givenIncludeDetails)
	}
}

func TestGetContactRequest_Execute(t *testing.T) {
	expectedEmail := test.RandomString(12, 36)
	var givenEmail string

	expectedResult := randomPerson()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmail = req.FormValue("email")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := contacts2.GetContact(req, expectedEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmail != givenEmail {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedEmail, givenEmail)
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatal("Results should be equal")
	}
}

func randomPerson() *contacts2.Person {
	var lists []contacts2.PersonList
	l := test.RandomInt(1, 12)
	for i := 0; i < l; i++ {
		lists = append(lists, contacts2.PersonList{
			ID:      test.RandomInt64(9999, 999999),
			Title:   test.RandomString(12, 36),
			AddedAt: test.RandomTime(12, 365),
		})
	}

	fields := map[string]string{}
	l = test.RandomInt(1, 12)
	for i := 0; i < l; i++ {
		fields[test.RandomString(4, 16)] = test.RandomString(12, 36)
	}

	return &contacts2.Person{
		Email: contacts2.PersonEmail{
			Email:        test.RandomString(12, 36),
			AddedAt:      test.RandomTime(12, 365),
			Status:       test.RandomString(12, 36),
			Availability: test.RandomString(12, 36),
			LastSend:     test.RandomTime(12, 365),
			LastDelivery: test.RandomTime(12, 365),
			LastRead:     test.RandomTime(12, 365),
			LastClick:    test.RandomTime(12, 365),
			Rating:       1.2,
			Lists:        lists,
			Fields:       fields,
		},
	}
}
